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
		mods := viper.GetStringSlice("mods")
		version, err := cmd.PersistentFlags().GetString("f-version")
		if err != nil {
			log.WithError(err).Warn("Failed to parse --f-version")
		}

		if version == "" {
			v, err := pkg.GetFactorioVersion()
			if err != nil {
				log.WithError(err).Fatal("Unknow Factorio version")
			}
			version = v.GetModString()
		}
		log.Infof("Factorio version is %s", version)

		// Download our mods, only warn if it could not be downloaded.
		for _, mod := range mods {
			log.Infof("Finding latest version of %s for Factorio %s", mod, version)
			m := pkg.GetReleaseMeta(mod)
			r, err := m.GetLatestReleaseByVersion(version)
			if err != nil {
				log.WithError(err).Warnf("Failed to find compatible version for %s and Factorio %s", mod, version)
			} else {
				err := pkg.DownloadModRelease(r)
				if err != nil {
					log.WithError(err).Warnf("Failed to download %s", r.FileName)
				}
				log.Infof("Downloaded %s (%s)", m.Name, r.Version)
			}
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
