package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/mdlv"
	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/rest/handlers"
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

type Rest struct {
	router   *chi.Mux
	server   *http.Server
	handlers Handlers

	log *logrus.Entry
	cfg config.Config
}

func NewRest(cfg config.Config, log *logrus.Logger, app *app.App) Rest {
	logger := log.WithField("module", "rest")

	router := chi.NewRouter()
	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}

	hands := handlers.NewHandlers(cfg, logger, app)

	return Rest{
		router:   router,
		server:   server,
		handlers: hands,

		log: logger,
		cfg: cfg,
	}
}

func (a *Rest) Run(ctx context.Context, log *logrus.Logger) {
	auth := mdlv.AuthMdl(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey)
	adminGrant := mdlv.AccessGrant(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey, roles.Admin, roles.SuperUser)

	a.router.Route("/chains/media-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/media", func(r chi.Router) {
				r.With(auth).Post("/", a.handlers.UploadMedia)
				r.Route("/{media_id}", func(r chi.Router) {
					r.Get("/", a.handlers.GetMedia)
					r.With(auth).Delete("/", a.handlers.DeleteMedia)
				})
			})

			r.Route("/media-rules", func(r chi.Router) {
				r.Route("/{resource}", func(r chi.Router) {
					r.With(adminGrant).Post("/", a.handlers.CreateMediaRules)
					r.With(adminGrant).Patch("/", a.handlers.UpdateMediaRules)
					r.With(adminGrant).Delete("/", a.handlers.DeleteMediaRules)
					r.Get("/", a.handlers.GetMediaRules)
				})
			})
		})
	})

	a.Start(ctx, log)

	<-ctx.Done()
	a.Stop(ctx, log)
}

func (a *Rest) Start(ctx context.Context, log *logrus.Logger) {
	go func() {
		a.log.Infof("Starting server on port %s", a.cfg.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
}

func (a *Rest) Stop(ctx context.Context, log *logrus.Logger) {
	a.log.Info("Shutting down server...")
	if err := a.server.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown failed: %v", err)
	}
}
