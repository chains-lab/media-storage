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
)

// MediaModel описывает метаданные объекта в S3
type MediaModel struct {
	Filename string
	Folder   string // имя или относительный путь файла в бакете
	Ext      string
	Size     int64  // размер в байтах
	URL      string // публичный URL объекта
}

// S3Client обёртка над AWS SDK для работы с S3
type S3Client struct {
	client *s3.Client
	bucket string
	region string
}

func NewAwsS3Client(bucket, region, accessKeyID, secretAccessKey string) (*S3Client, error) {
	// Создаём провайдер, который отдаёт именно эти ключи
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

type AddFileInput struct {
	Reader io.Reader
}

// AddFile загружает файл в указанную папку (folder) с именем filename и возвращает публичный URL
func (s *S3Client) AddFile(ctx context.Context, folder, filename, ext string, input AddFileInput) (MediaModel, error) {
	key := path.Join(folder, filename+ext)

	ct := mime.TypeByExtension(ext)
	if ct == "" {
		return MediaModel{}, fmt.Errorf("unsupported file extension: %s", ext)
	}
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        input.Reader,
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
		Filename: filename,
		Folder:   folder,
		URL:      url,
		Size:     size,
	}, nil
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

			url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
			results = append(results, MediaModel{
				Filename: filename,
				Folder:   folder,
				URL:      url,
				Size:     aws.ToInt64(obj.Size),
			})
		}
	}
	return results, nil
}

func (s *S3Client) DeleteFile(ctx context.Context, folder, filename, ext string) error {
	name := path.Join(folder, filename+ext)
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return fmt.Errorf("delete S3 object %s: %w", name, err)
	}
	return nil
}

// DeleteFilesInFolder удаляет только непосредственные файлы (не папки) в указанном folder
func (s *S3Client) DeleteFilesInFolder(ctx context.Context, folder string) error {
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.bucket),
		Prefix:    aws.String(folder + "/"),
		Delimiter: aws.String("/"),
	}
	page, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return fmt.Errorf("listing objects for delete: %w", err)
	}
	if len(page.Contents) == 0 {
		return nil
	}

	var toDelete []s3types.ObjectIdentifier
	for _, obj := range page.Contents {
		toDelete = append(toDelete, s3types.ObjectIdentifier{Key: obj.Key})
	}
	_, err = s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(s.bucket),
		Delete: &s3types.Delete{Objects: toDelete, Quiet: aws.Bool(true)},
	})
	if err != nil {
		return fmt.Errorf("batch delete objects in %s: %w", folder, err)
	}
	return nil
}
