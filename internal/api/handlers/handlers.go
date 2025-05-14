package handlers

import (
	"context"

	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type App interface {
	GetMedia(ctx context.Context, mediaID uuid.UUID) (app.MediaModels, error)
	UploadMedia(ctx context.Context, request app.UploadMediaRequest) (app.MediaModels, error)
	DeleteMedia(ctx context.Context, request app.DeleteMediaRequest) error

	CreateMediaRules(ctx context.Context, request app.CreateMediaRulesRequest) (app.MediaRulesModel, error)
	GetMediaRules(ctx context.Context, id string) (app.MediaRulesModel, error)
	UpdateMediaRules(ctx context.Context, id string, request app.UpdateMediaRulesRequest) (app.MediaRulesModel, error)
	DeleteMediaRules(ctx context.Context, id string) error
}

type Handler struct {
	app App
	cfg config.Config
	log *logrus.Entry
}

func NewHandlers(cfg config.Config, log *logrus.Entry, app *app.App) Handler {
	return Handler{
		app: app,
		cfg: cfg,
		log: log,
	}
}
