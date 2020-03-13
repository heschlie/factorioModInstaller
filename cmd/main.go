package cmd

import "github.com/spf13/cobra"

func init() {
	
}

var rootCmd = &cobra.Command{
	Use: "mod-install",
	Short: "Installs and updates Factorio mods for headless servers",
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a list of mods",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
