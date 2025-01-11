package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile  string
	dotfilesDir string

	rootCmd = &cobra.Command{
		Use:   "dotx",
		Short: "The next generation dotfile manager",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default $HOME/.dotx.toml)")
	rootCmd.PersistentFlags().StringVarP(&dotfilesDir, "dir", "d", "$HOME/.dotfiles", "dotfiles directory")

	viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)

		viper.SetConfigType("toml")
		viper.SetConfigName(".dotx")
	}

	viper.AutomaticEnv()

	viper.ReadInConfig()
}
