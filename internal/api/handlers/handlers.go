package handlers

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/tokens"
	"github.com/sirupsen/logrus"
)

type App interface {
	UploadMedia(ctx context.Context, user tokens.AccountData, file io.Reader, fileHeader *multipart.FileHeader, request app.UploadMediaRequest) (models.Media, error)
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
