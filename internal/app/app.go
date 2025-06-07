package app

import (
	"context"

	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/domain"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/google/uuid"
)

type media interface {
	Upload(ctx context.Context, request domain.UploadMediaRequest) (models.Media, *ape.Error)
	Get(ctx context.Context, mediaID uuid.UUID) (models.Media, *ape.Error)
	Delete(ctx context.Context, request domain.DeleteMediaRequest) *ape.Error
}

type mediaRules interface {
	Get(ctx context.Context, ID string) (models.MediaRules, *ape.Error)
	Create(ctx context.Context, request domain.CreateMediaRulesRequest) (models.MediaRules, *ape.Error)
	Update(ctx context.Context, ID string, request domain.UpdateMediaRulesRequest) (models.MediaRules, *ape.Error)
	Delete(ctx context.Context, ID string) *ape.Error
}

type App struct {
	media media
	rules mediaRules
}

func NewApp(cfg config.Config) (App, error) {
	mediaDomain, err := domain.NewMedia(cfg)
	if err != nil {
		return App{}, err
	}

	rulesDomain, err := domain.NewMediaRules(cfg)
	if err != nil {
		return App{}, err
	}

	return App{
		media: mediaDomain,
		rules: rulesDomain,
	}, nil
}
