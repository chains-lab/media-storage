package app

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/repo"
)

type repoMedia interface {
	UploadMedia(ctx context.Context, reader io.Reader, input repo.AddMediaInput) (repo.MediaModel, error)
	GetMedia(ctx context.Context, mediaID uuid.UUID) (repo.MediaModel, error)
	DeleteMedia(ctx context.Context, mediaID uuid.UUID) error
	DeleteFilesByResourceAndCategory(ctx context.Context, resource, category string) error
}

type MediaRulesRepo interface {
	Create(ctx context.Context, input repo.CreateMediaRulesInput) (repo.MediaRulesModel, error)
	Get(ctx context.Context, id string) (repo.MediaRulesModel, error)
	Update(ctx context.Context, id string, input repo.MediaRulesUpdateInput) error
	Delete(ctx context.Context, id string) error
}

type App struct {
	repoMedia repoMedia
	repoRules MediaRulesRepo
}

func NewApp(cfg config.Config) (App, error) {
	mediaRepo, err := repo.NewMedia(cfg)
	if err != nil {
		return App{}, err
	}

	rulesRepo, err := repo.NewMediaRulesRepo(cfg)
	if err != nil {
		return App{}, err
	}

	return App{
		repoMedia: mediaRepo,
		repoRules: rulesRepo,
	}, nil
}
