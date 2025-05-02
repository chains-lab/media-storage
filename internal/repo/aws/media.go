package aws

import (
	"context"
	"fmt"
	"io"
	"mime"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type MediaModel struct {
	Filename uuid.UUID
	Folder   string
	Ext      string
	Size     int64
	URL      string
}

type FileData struct {
	Filename uuid.UUID
	Folder   string
	Ext      string
}

type S3Client struct {
	client *s3.Client
	bucket string
	region string
}

func NewAwsS3Client(bucket, region, accessKeyID, secretAccessKey string) (*S3Client, error) {
	staticProvider := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(),
		awscfg.WithRegion(region),
		awscfg.WithCredentialsProvider(staticProvider),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
		bucket: bucket,
		region: region,
	}, nil
}

func (s *S3Client) AddFile(ctx context.Context, fileData FileData, reader io.Reader) (MediaModel, error) {
	key := path.Join(fileData.Folder, fileData.Filename.String()+fileData.Ext)

	ct := mime.TypeByExtension(fileData.Ext)
	if ct == "" {
		return MediaModel{}, fmt.Errorf("unsupported file extension: %s", fileData.Ext)
	}
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(ct),
	})

	if err != nil {
		return MediaModel{}, fmt.Errorf("put S3 object %s: %w", key, err)
	}

	head, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return MediaModel{}, fmt.Errorf("head S3 object %s: %w", key, err)
	}

	size := aws.ToInt64(head.ContentLength)
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
	return MediaModel{
		Filename: fileData.Filename,
		Folder:   fileData.Folder,
		URL:      url,
		Size:     size,
	}, nil
}

func (s *S3Client) GetFile(ctx context.Context, fileData FileData) (MediaModel, error) {
	key := path.Join(fileData.Folder, fileData.Filename.String()+fileData.Ext)
	head, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return MediaModel{}, fmt.Errorf("head S3 object %s: %w", key, err)
	}

	size := aws.ToInt64(head.ContentLength)
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
	return MediaModel{
		Filename: fileData.Filename,
		Folder:   fileData.Folder,
		Ext:      fileData.Ext,
		URL:      url,
		Size:     size,
	}, nil
}

func (s *S3Client) DeleteFile(ctx context.Context, fileData FileData) error {
	name := path.Join(fileData.Folder, fileData.Filename.String()+fileData.Ext)
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return fmt.Errorf("delete S3 object %s: %w", name, err)
	}
	return nil
}

func (s *S3Client) ListFiles(ctx context.Context, folder string, offset, limit uint) ([]MediaModel, error) {
	var results []MediaModel
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(folder + "/"),
	}
	paginator := s3.NewListObjectsV2Paginator(s.client, input)
	count := uint(0)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("listing S3 objects: %w", err)
		}
		for _, obj := range page.Contents {
			if count < offset {
				count++
				continue
			}

			if uint(len(results)) >= limit {
				return results, nil
			}

			key := aws.ToString(obj.Key)
			parts := strings.Split(key, "/")

			var filename string
			if len(parts) > 1 {
				filename = parts[len(parts)-1]
			} else {
				filename = key
			}

			parts = strings.Split(filename, ".")
			var ext string
			var fileID uuid.UUID
			if len(parts) > 1 {
				ext = parts[len(parts)-1]
				fileID, err = uuid.Parse(parts[0])
				if err != nil {
					return nil, fmt.Errorf("parsing UUID from filename %s: %w", filename, err)
				}
			} else {
				ext = ""
			}

			url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
			results = append(results, MediaModel{
				Filename: fileID,
				Folder:   folder,
				Ext:      ext,
				URL:      url,
				Size:     aws.ToInt64(obj.Size),
			})
		}
	}
	return results, nil
}

func (s *S3Client) DeleteFilesByPrefix(ctx context.Context, prefix string) error {
	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("listing S3 objects with prefix %q: %w", prefix, err)
		}

		if len(page.Contents) == 0 {
			// больше нет объектов под этим префиксом
			return nil
		}

		// Собираем идентификаторы для батч-удаления
		identifiers := make([]s3types.ObjectIdentifier, 0, len(page.Contents))
		for _, obj := range page.Contents {
			identifiers = append(identifiers, s3types.ObjectIdentifier{
				Key: obj.Key,
			})
		}

		// Удаляем до 1000 штук за раз
		quite := true
		_, err = s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(s.bucket),
			Delete: &s3types.Delete{
				Objects: identifiers,
				Quiet:   &quite,
			},
		})
		if err != nil {
			return fmt.Errorf("deleting S3 objects with prefix %q: %w", prefix, err)
		}
	}

	return nil
}
