package main

import (
	"factorioModInstaller/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.WithError(err).Error("Encountered error running command")
	}
}
