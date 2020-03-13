package cmd

import (
	"factorioModInstaller/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var configPath string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configPath, "config",
		filepath.Join("opt", "factorio", "mod-install.yaml"), "config file, default is /opt/factorio/mod-install.yaml")
	rootCmd.PersistentFlags().StringP("f-version", "f", "", "force Factorio version")
}

var rootCmd = &cobra.Command{
	Use: "mod-install",
	Short: "Installs and updates Factorio mods for headless servers",
	Long: "Installs or upgrades all mods listed in the mods section of the config. It will try to verify the Factorio" +
		"version by checking the binary, alternatively you can pass it in with --f-version 0.15",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.PersistentFlags().GetString("f-version")
		//fPath := viper.GetString("factorioPath")
		if err != nil {
			log.WithError(err).Warn("Failed to parse --f-version")
		}
		if version == "" {

		}

		mods := viper.GetStringSlice("mods")
		for _, mod := range mods {
			log.Infof("Finding latest version of %s for Factorio %s", mod, )
			m := pkg.GetReleaseMeta(mod)
			r, _ := m.GetLatestReleaseByVersion("0.17")
			log.Infof("%s: %s", m.Name, r.DownloadURL)
		}
	},
}

func initConfig() {
	viper.SetConfigFile(configPath)
	viper.SetDefault("mods", []string{})
	viper.SetDefault("factorioPath", filepath.Join("opt", "factorio"))

	err := viper.ReadInConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to find config")
	}
}
