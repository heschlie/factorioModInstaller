package main

import (
	"factorioModInstaller/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
)

func main() {
	//err := cmd.Execute()
	//if err != nil {
	//	log.WithError(err).Error("Encountered error running command")
	//}
	viper.SetDefault("mods", []string{})
	viper.SetDefault("factorioPath", filepath.Join("/home/heschlie/Downloads/factorio_headless_x64_0.17.79/factorio"))

	viper.ReadInConfig()

	v, _ := pkg.GetFactorioVersion()
	log.Info(v.String())
}
