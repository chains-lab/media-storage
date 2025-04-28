package app

import (
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/sirupsen/logrus"
)

type App struct {
}

func NewApp(cfg config.Config, log *logrus.Logger) (App, error) {
	return App{}, nil
}
