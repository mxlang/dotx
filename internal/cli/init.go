package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdInit(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Initialize by cloning a remote dotfiles repository",
		Long:    "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files",
		Example: `  dotx sync init https://github.com/username/dotfiles.git`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.CloneRemoteRepo(args[0]); err != nil {
				logger.Error("failed to clone remote repo", "error", err)
			}

			logger.Info("successfully cloned remote repo")
		},
	}
}
