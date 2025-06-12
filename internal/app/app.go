package app

import (
	"context"

	"github.com/chains-lab/media-storage/internal/app/domain"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/aws"
	"github.com/chains-lab/media-storage/internal/config"
)

type mediaRules interface {
	GetMediaRules(ctx context.Context, resource, category string) (models.MediaRules, error)
	GenerateMediaPutURL(ctx context.Context, folder, originalFilename, contentType string) (*aws.MediaURL, error)
}

type App struct {
	rules mediaRules
}

func NewApp(cfg config.Config) (App, error) {
	rulesDomain, err := domain.NewMediaRules(cfg)
	if err != nil {
		return App{}, err
	}

	return App{
		rules: rulesDomain,
	}, nil
}
