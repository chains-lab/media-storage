package aws

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ContentModel описывает метаданные объекта в S3
// Key — ключ (путь), URL — полный адрес, Size — размер, LastModified — время изменений
// Эти данные могут быть использованы для любых целей в бизнес-логике
// (отображение, логирование, аналитика и т.д.)
type ContentModel struct {
	Key          string    // Ключ объекта в бакете
	URL          string    // Полный URL для доступа
	Size         int64     // Размер в байтах
	LastModified time.Time // Дата последнего изменения
}

// S3Client реализует базовые операции над S3 без бизнес-логики
// Это инфраструктурный адаптер, который умеет: list, get presigned URL, put, delete, batch delete
// Все эти методы возвращают простые структуры и ошибки fmt.Errorf, чтобы на уровне application можно было их обрабатывать
type S3Client struct {
	client    *s3.Client
	presigner *s3.PresignClient
	bucket    string
	region    string
}

// NewS3Client создаёт новый экземпляр клиента для работы с S3
// Читает конфиг из окружения, создаёт s3.Client и s3.PresignClient
func NewS3Client(ctx context.Context, bucket, region string) (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("loading AWS config: %w", err)
	}
	client := s3.NewFromConfig(cfg)
	presigner := s3.NewPresignClient(client)
	return &S3Client{
		client:    client,
		presigner: presigner,
		bucket:    bucket,
		region:    region,
	}, nil
}

// GetContent возвращает список объектов под указанным префиксом (path)
func (s *S3Client) GetContent(ctx context.Context, path string) ([]ContentModel, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(path),
	}
	out, err := s.client.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("listing S3 objects: %w", err)
	}

	var result []ContentModel
	for _, obj := range out.Contents {
		key := aws.ToString(obj.Key)
		url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
		result = append(result, ContentModel{
			Key:          key,
			URL:          url,
			Size:         *obj.Size,
			LastModified: aws.ToTime(obj.LastModified),
		})
	}
	return result, nil
}

// AddContent загружает содержимое reader в S3 по указанному ключу (key)
// и возвращает полный URL для доступа к файлу
func (s *S3Client) AddContent(ctx context.Context, key string, reader io.Reader, contentType string) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	}
	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return "", fmt.Errorf("put S3 object: %w", err)
	}
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
	return url, nil
}

// DeleteContent удаляет один объект по ключу (key)
func (s *S3Client) DeleteContent(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		return fmt.Errorf("delete S3 object: %w", err)
	}
	return nil
}

// CleanContent удаляет все объекты под указанным префиксом (path)
func (s *S3Client) CleanContent(ctx context.Context, path string) error {
	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(path),
	}
	listOut, err := s.client.ListObjectsV2(ctx, listInput)
	if err != nil {
		return fmt.Errorf("listing objects for cleanup: %w", err)
	}
	if len(listOut.Contents) == 0 {
		return nil
	}

	// Формируем список идентификаторов
	var identifiers []types.ObjectIdentifier
	for _, obj := range listOut.Contents {
		identifiers = append(identifiers, types.ObjectIdentifier{Key: obj.Key})
	}

	// Пакетное удаление
	delInput := &s3.DeleteObjectsInput{
		Bucket: aws.String(s.bucket),
		Delete: &types.Delete{
			Objects: identifiers,
			Quiet:   aws.Bool(true),
		},
	}
	_, err = s.client.DeleteObjects(ctx, delInput)
	if err != nil {
		return fmt.Errorf("batch delete objects: %w", err)
	}
	return nil
}

// PresignUpload возвращает presigned URL для загрузки PUT
// expires задает время жизни ссылки
func (s *S3Client) PresignUpload(ctx context.Context, key, contentType string, expires time.Duration) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}
	presignResult, err := s.presigner.PresignPutObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = expires
	})
	if err != nil {
		return "", fmt.Errorf("presign upload URL: %w", err)
	}
	return presignResult.URL, nil
}
