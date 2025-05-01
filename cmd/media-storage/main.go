package main

import (
	"os"

	"github.com/hs-zavet/media-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
