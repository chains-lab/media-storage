package domain

import (
	"context"

	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/aws"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo"
	"github.com/chains-lab/media-storage/internal/repo/sqldb"
)

type mediaRulesRepo interface {
	GetByResourceAndCategory(ctx context.Context, resource, category string) (sqldb.MediaRulesModel, error)
}

type MediaAllowedExtensionRepo interface {
	GetByResourcesAndCategory(ctx context.Context, resource, category string) ([]sqldb.AllowedExtensionModel, error)
}

type awsS3 interface {
	GeneratePutURL(ctx context.Context, folder, originalFilename, contentType string) (*aws.MediaURL, error)
	GenerateGetURL(ctx context.Context, key string) (string, error)
}
type MediaRules struct {
	rules      mediaRulesRepo
	extensions MediaAllowedExtensionRepo
	s3         awsS3
}

func NewMediaRules(cfg config.Config) MediaRules {
	rulesRepo, err := repo.NewMediaRules(cfg)
	if err != nil {
		panic(err)
	}

	extensionsRepo, err := repo.NewAllowedExtension(cfg)
	if err != nil {
		panic(err)
	}

	return MediaRules{
		rules:      rulesRepo,
		extensions: extensionsRepo,
	}

}

func (m MediaRules) GetMediaRules(ctx context.Context, resource, category string) (models.MediaRules, error) {
	mediaRules, err := m.rules.GetByResourceAndCategory(ctx, resource, category)
	if err != nil {
		return models.MediaRules{}, err
	}

	allowedExt, err := m.extensions.GetByResourcesAndCategory(ctx, resource, category)
	if err != nil {
		return models.MediaRules{}, err
	}

	extensions := make([]models.AllowedExtensionModel, 0, len(allowedExt))
	for _, ext := range allowedExt {
		extensions = append(extensions, models.AllowedExtensionModel{
			Extension: ext.Extension,
			MaxSize:   ext.MaxSize,
		})
	}

	return models.MediaRules{
		Resource:   mediaRules.Resource,
		Category:   mediaRules.Category,
		Extensions: extensions,
	}, nil
}

func (m MediaRules) GenerateMediaPutURL(ctx context.Context, folder, originalFilename, contentType string) (*aws.MediaURL, error) {
	putURL, err := m.s3.GeneratePutURL(ctx, folder, originalFilename, contentType)
	if err != nil {
		return nil, err
	}

	return putURL, nil
}
