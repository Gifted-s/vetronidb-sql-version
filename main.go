package main

import (
	"os"
	"time"

	"github.com/Gifted-s/sunkdb/cli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
func main(){
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	cli.StartPrompt("random file")
}