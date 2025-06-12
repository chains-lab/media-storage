package cli

import (
	"context"
	"sync"

	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/rest"
	"github.com/sirupsen/logrus"
)

func runServices(ctx context.Context, cfg config.Config, log *logrus.Logger, wg *sync.WaitGroup, app *app.App) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	API := rest.NewRest(cfg, log, app)
	run(func() { API.Run(ctx, log) })
}
