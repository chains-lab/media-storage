package handlers

import (
	"github.com/chains-lab/media-storage/internal/api/controllers"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/domain"
	"github.com/sirupsen/logrus"
)

type Controllers interface {
}

type Handler struct {
	controllers Controllers
	cfg         config.Config
	log         *logrus.Logger
}

func NewHandlers(cfg config.Config, log *logrus.Logger, app *domain.App) Handler {
	cntrl := controllers.NewController(log, cfg, app)

	return Handler{
		cfg: cfg,
		log: log,
	}
}
