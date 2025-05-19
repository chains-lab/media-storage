package handlers

import (
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	app    *app.App
	log    *logrus.Logger
}

func NewHandlers(log *logrus.Logger, app *app.App) Handler {
return Handler{
		app:    app,
		log:    log,
	}
}
