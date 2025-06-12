package aws

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

// S3Client handles generating presigned URLs instead of direct uploads
type S3Client struct {
	client     *s3.Client
	presignCli *s3.PresignClient
	bucket     string
	region     string
	defaultTTL time.Duration
}

// MediaURL holds data about the upload URL and object metadata
type MediaURL struct {
	Key       string        // S3 key
	UploadURL string        // presigned PUT URL
	ExpiresIn time.Duration // TTL of the URL
}

// NewS3PresignClient initializes the S3 client and presign client
func NewS3PresignClient(bucket, region, accessKeyID, secretAccessKey string, defaultTTL time.Duration) (*S3Client, error) {
	staticProvider := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(),
		awscfg.WithRegion(region),
		awscfg.WithCredentialsProvider(staticProvider),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	cli := s3.NewFromConfig(cfg)
	ps := s3.NewPresignClient(cli)
	return &S3Client{
		client:     cli,
		presignCli: ps,
		bucket:     bucket,
		region:     region,
		defaultTTL: defaultTTL,
	}, nil
}

// GeneratePutURL returns a presigned PUT URL for uploading a file
func (s *S3Client) GeneratePutURL(ctx context.Context, folder, originalFilename, contentType string) (*MediaURL, error) {
	// Extract extension and generate unique filename
	ext := path.Ext(originalFilename)
	id := uuid.New().String()
	key := path.Join(folder, id+ext)

	// Prepare presign parameters
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}

	// Generate presigned URL
	resp, err := s.presignCli.PresignPutObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = s.defaultTTL
	})
	if err != nil {
		return nil, fmt.Errorf("failed to presign PUT object %s: %w", key, err)
	}

	return &MediaURL{
		Key:       key,
		UploadURL: resp.URL,
		ExpiresIn: s.defaultTTL,
	}, nil
}

// GenerateGetURL returns a presigned GET URL for retrieving a file
func (s *S3Client) GenerateGetURL(ctx context.Context, key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}
	resp, err := s.presignCli.PresignGetObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = s.defaultTTL
	})
	if err != nil {
		return "", fmt.Errorf("failed to presign GET object %s: %w", key, err)
	}
	return resp.URL, nil
}
