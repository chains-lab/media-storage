package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/sirupsen/logrus"
)

type App interface {
	GetMedia(ctx context.Context, resourceID uuid.UUID) (models.Media, error)
	UploadMedia(ctx context.Context, request app.UploadMediaRequest) (models.Media, error)
	DeleteMedia(ctx context.Context, request app.DeleteMediaRequest) error

	CreateMediaRules(ctx context.Context, request app.CreateMediaRulesRequest) (models.MediaRules, error)
	GetMediaRules(ctx context.Context, mediaType enums.MediaType) (models.MediaRules, error)
	UpdateMediaRules(ctx context.Context, mType enums.MediaType, request app.UpdateMediaRulesRequest) (models.MediaRules, error)
	DeleteMediaRules(ctx context.Context, mType enums.MediaType) error
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
