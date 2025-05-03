package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hs-zavet/media-storage/internal/api/handlers"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/tokens"
	"github.com/hs-zavet/tokens/roles"
	"github.com/sirupsen/logrus"
)

type Api struct {
	server   *http.Server
	router   *chi.Mux
	handlers handlers.Handler

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
		server:   server,
		router:   router,
		handlers: hands,

		log: logger,
		cfg: cfg,
	}
}

func (a *Api) Run(ctx context.Context, log *logrus.Logger) {
	auth := tokens.AuthMdl(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey)
	adminGrant := tokens.AccessGrant(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey, roles.Admin, roles.SuperUser)

	a.router.Route("/hs-news/media-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/media", func(r chi.Router) {
				r.Route("/{resource_type}", func(r chi.Router) {
					r.With(auth).Post("/", a.handlers.UploadMedia)
					r.With(adminGrant).Delete("/{media_id}", a.handlers.DeleteMedia)
					r.Get("/{media_id}", a.handlers.GetMedia)
				})
			})

			r.Route("/media-rules", func(r chi.Router) {
				r.Route("/{resource_type}", func(r chi.Router) {
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

func (a *Api) Start(ctx context.Context, log *logrus.Logger) {
	go func() {
		a.log.Infof("Starting server on port %s", a.cfg.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
}

func (a *Api) Stop(ctx context.Context, log *logrus.Logger) {
	a.log.Info("Shutting down server...")
	if err := a.server.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown failed: %v", err)
	}
}
