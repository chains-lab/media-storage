package aws

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	s3Manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type MediaModel struct {
	ID          string
	Extension   string
	ContentType string
	Folder      string
	Size        int64
	Url         string
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

// UploadFile uploads a file to S3 and returns a MediaModel with the file information.
// filename is ORIGINAL filename, without changes
func (s *S3Client) UploadFile(ctx context.Context, reader io.Reader, folder, filename string) (MediaModel, error) {
	buf := make([]byte, 512)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		return MediaModel{}, fmt.Errorf("read header: %w", err)
	}
	ct := http.DetectContentType(buf[:n])

	extOrig := ""
	if idx := strings.Index(filename, "."); idx != -1 {
		extOrig = filename[idx:]
		filename = uuid.New().String()
	}

	var ext string
	if extOrig != "" {
		ext = extOrig
	} else {
		ext = ""
	}

	reader = io.MultiReader(bytes.NewReader(buf[:n]), reader)

	key := path.Join(folder, filename+ext)
	uploader := s3Manager.NewUploader(s.client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
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
	urlRes := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
	return MediaModel{
		ID:          filename,
		Extension:   ext,
		ContentType: ct,
		Folder:      folder,
		Url:         urlRes,
		Size:        size,
	}, nil
}

func (s *S3Client) GetFileByUrl(ctx context.Context, rawURL string) (MediaModel, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return MediaModel{}, fmt.Errorf("error parse Url %q: %w", rawURL, err)
	}
	key := strings.TrimPrefix(u.Path, "/")

	head, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return MediaModel{}, fmt.Errorf("head S3 object %q: %w", key, err)
	}

	var folder string
	parts := strings.Split(key, "/")
	if len(parts) > 1 {
		folder = strings.Join(parts[:len(parts)-1], "/")
	} else {
		folder = ""
	}

	var ct string
	if head.ContentType != nil {
		ct = *head.ContentType
	} else {
		ct = mime.TypeByExtension(path.Ext(key))
	}

	filename := parts[len(parts)-1]
	ext := path.Ext(filename)

	size := aws.ToInt64(head.ContentLength)
	return MediaModel{
		ID:          filename,
		Extension:   ext,
		ContentType: ct,
		Folder:      folder,
		Size:        size,
		Url:         rawURL,
	}, nil
}

func (s *S3Client) DeleteFileByUrl(ctx context.Context, rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("error parse Url %q: %w", rawURL, err)
	}
	key := strings.TrimPrefix(u.Path, "/")

	_, err = s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("delete S3 object %s: %w", key, err)
	}
	return nil
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
