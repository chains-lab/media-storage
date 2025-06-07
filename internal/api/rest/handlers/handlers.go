package handlers

import (
	"context"
	"net/http"

	"github.com/chains-lab/media-storage/internal/api/rest/presenter"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type App interface {
	UploadMedia(ctx context.Context, request app.UploadMediaRequest) (models.Media, *ape.Error)
	GetMedia(ctx context.Context, mediaID uuid.UUID) (models.Media, *ape.Error)
	DeleteMedia(ctx context.Context, request app.DeleteMediaRequest) *ape.Error

	CreateMediaRules(ctx context.Context, request app.CreateMediaRulesRequest) (models.MediaRules, *ape.Error)
	GetMediaRules(ctx context.Context, ID string) (models.MediaRules, *ape.Error)
	UpdateMediaRules(ctx context.Context, ID string, request app.UpdateMediaRulesRequest) (models.MediaRules, *ape.Error)
	DeleteMediaRules(ctx context.Context, ID string) *ape.Error
}

type Presenter interface {
	InvalidPointer(w http.ResponseWriter, requestID uuid.UUID, err error)
	InvalidToken(w http.ResponseWriter, requestID uuid.UUID, err error)
	InvalidParameter(w http.ResponseWriter, requestID uuid.UUID, err error, parameter string)
	InvalidQuery(w http.ResponseWriter, requestID uuid.UUID, query string, err error)
	MismatchIdentification(w http.ResponseWriter, requestID uuid.UUID, parameter, pointer string)
	AppError(w http.ResponseWriter, requestID uuid.UUID, appErr *ape.Error)
}

type Handler struct {
	app       App
	presenter Presenter
	log       *logrus.Entry
	cfg       config.Config
}

func NewHandlers(cfg config.Config, log *logrus.Entry, app *app.App) Handler {
	return Handler{
		app:       app,
		cfg:       cfg,
		presenter: presenter.NewPresenter(log),
		log:       log,
	}
}
