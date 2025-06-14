package api

import (
	"net/http"

	"github.com/chains-lab/media-storage/internal/api/rest/handlers"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Handlers interface {
	UploadMedia(w http.ResponseWriter, r *http.Request)
	GetMedia(w http.ResponseWriter, r *http.Request)
	DeleteMedia(w http.ResponseWriter, r *http.Request)

	CreateMediaRules(w http.ResponseWriter, r *http.Request)
	UpdateMediaRules(w http.ResponseWriter, r *http.Request)
	GetMediaRules(w http.ResponseWriter, r *http.Request)
	DeleteMediaRules(w http.ResponseWriter, r *http.Request)
}

type Api struct {
	router   *chi.Mux
	server   *http.Server
	handlers Handlers

	log *logrus.Entry
	cfg config.Config
}

func NewAPI(cfg config.Config, log *logrus.Logger, app *app.App) Api {
	logger := log.WithField("module", "api")

	router := chi.NewRouter()
	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}

	hands := handlers.NewHandlers(cfg, logger, app)

	return Api{
		router:   router,
		server:   server,
		handlers: hands,

		log: logger,
		cfg: cfg,
	}
}
