package cli

import (
	"github.com/mxlang/dotx/internal/logger"
	"os"

	"github.com/mxlang/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func Execute(dotx dotx.App) {
	err := newCmdRoot(dotx).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(dotx dotx.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dotx",
		Short: "A modern dotfile manager for tracking and syncing configuration files",
		Long:  "dotx helps you manage, version control, and synchronize your configuration files (dotfiles) across multiple systems",
		Args:  cobra.NoArgs,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if dotx.AppConfig.Verbose {
				logger.SetLevel(logger.DebugLevel)
			}
		},
	}

	rootCmd.AddCommand(
		newCmdVersion(dotx),
		newCmdAdd(dotx),
		newCmdDeploy(dotx),
		newCmdSync(dotx),
	)

	return rootCmd
}
