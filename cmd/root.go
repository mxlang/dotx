package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dotx",
	Short: "The next generation dotfile manager",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("dir", "d", "$HOME/.dotfiles", "dotfiles directory")
	viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigName(".dotx")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(home)

	viper.ReadInConfig()
}
