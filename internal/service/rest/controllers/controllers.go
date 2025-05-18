package controllers

import (
	"context"

	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type App interface {
	GetMedia(ctx context.Context, mediaID uuid.UUID) (domain.MediaModels, error)
	UploadMedia(ctx context.Context, request domain.UploadMediaRequest) (domain.MediaModels, error)
	DeleteMedia(ctx context.Context, request domain.DeleteMediaRequest) error

	CreateMediaRules(ctx context.Context, request domain.CreateMediaRulesRequest) (domain.MediaRulesModel, error)
	GetMediaRules(ctx context.Context, id string) (domain.MediaRulesModel, error)
	UpdateMediaRules(ctx context.Context, id string, request domain.UpdateMediaRulesRequest) (domain.MediaRulesModel, error)
	DeleteMediaRules(ctx context.Context, id string) error
}

type Controller struct {
	app App

	log *logrus.Logger
	cfg config.Config
}

func NewController(cfg *config.Config, log *logrus.Logger, app *domain.App) Controller {
	return Controller{
		app: app,
		cfg: *cfg,
		log: log,
	}
}
