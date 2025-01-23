package cmd

import (
	"log/slog"
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

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
		},
	}

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "more verbose output")
	rootCmd.PersistentFlags().StringP("dir", "d", "$HOME/.dotfiles", "dotfiles directory on your machine")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))

	rootCmd.AddCommand(newInitCmd())

	return rootCmd
}
