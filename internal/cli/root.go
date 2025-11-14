package cli

import (
	"os"

	"github.com/mxlang/dotx/internal/cli/sync"
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/logger"

	"github.com/spf13/cobra"
)

func Execute(app core.App, version string) {
	err := newCmdRoot(app, version).Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newCmdRoot(app core.App, version string) *cobra.Command {
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

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "enable verbose output")

	rootCmd.AddCommand(
		newCmdInitShell(),

		// commands
		newCmdAdd(app),
		newCmdDeploy(app),
		newCmdCd(app),
		newCmdDoctor(app),

		// subcommand sync
		sync.NewCmdSync(app),
	)

	return rootCmd
}
