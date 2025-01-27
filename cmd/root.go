package cmd

import (
	"os"

	"github.com/mlang97/dotx/config"
	"github.com/spf13/cobra"
)

func Execute() {
	cfg := config.Load()

	err := newCmdRoot(cfg).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(cfg config.Config) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotx",
		Short: "The next generation dotfile manager",

		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 	if viper.GetBool("verbose") {
		// 		slog.SetLogLoggerLevel(slog.LevelDebug)
		// 	}
		// },
	}

	// rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "more verbose output")
	// rootCmd.PersistentFlags().StringVarP(&config.DotfilesDir, "dir", "d", "$HOME/.dotfiles", "dotfiles directory on your machine")

	// viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	// viper.BindPFlag("dir", rootCmd.PersistentFlags().Lookup("dir"))

	rootCmd.AddCommand(
		newCmdInit(cfg),
		newCmdAdd(cfg))

	return rootCmd
}
