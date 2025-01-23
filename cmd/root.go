package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	err := newCmdRoot().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotx",
		Short: "The next generation dotfile manager",
	}

	rootCmd.PersistentFlags().StringP("dir", "d", "$HOME/.dotfiles", "dotfiles directory")
	viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))

	rootCmd.AddCommand(newInitCmd())

	return rootCmd
}
