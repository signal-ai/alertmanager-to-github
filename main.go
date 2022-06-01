package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/signal-ai/alertmanager-to-github/pkg/cli"
)

func main() {
	err := cli.App().Run(os.Args)
	if err != nil {
		log.Fatal().Err(err)
	}
}
