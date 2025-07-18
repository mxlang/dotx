package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

func Execute(cfg *config.Config, version string) {
	err := newCmdRoot(cfg, version).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(cfg *config.Config, version string) *cobra.Command {
	var verbose bool

	rootCmd := &cobra.Command{
		Use:     "dotx",
		Short:   "A modern dotfile manager for tracking and syncing configuration files",
		Long:    "dotx helps you manage, version control, and synchronize your configuration files (dotfiles) across multiple systems",
		Version: version,

		Args: cobra.NoArgs,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				logger.SetLevel(logger.DebugLevel)
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", cfg.App.Verbose, "enable verbose output")

	rootCmd.AddCommand(
		newCmdAdd(cfg),
		newCmdDeploy(cfg),
		newCmdSync(cfg),
		newCmdCd(cfg),
		newCmdInitShell(),
	)

	return rootCmd
}
