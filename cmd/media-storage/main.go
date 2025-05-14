package main

import (
	"os"

	"github.com/chains-lab/media-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
