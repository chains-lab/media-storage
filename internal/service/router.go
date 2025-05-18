package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/mdlv"
	"github.com/chains-lab/gatekit/roles"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func (a *Api) Run(ctx context.Context, log *logrus.Logger) {
	auth := mdlv.AuthMdl(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey)
	adminGrant := mdlv.AccessGrant(a.cfg.JWT.AccessToken.SecretKey, a.cfg.JWT.ServiceToken.SecretKey, roles.Admin, roles.SuperUser)

	a.router.Route("/hs-news/media-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/media", func(r chi.Router) {
				r.With(auth).Post("/", a.handlers.UploadMedia)
				r.Route("/{media_id}", func(r chi.Router) {
					r.Get("/", a.handlers.GetMedia)
					r.With(auth).Delete("/", a.handlers.DeleteMedia)
				})
			})

			r.Route("/media-rules", func(r chi.Router) {
				r.Route("/{resource-category}", func(r chi.Router) {
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
