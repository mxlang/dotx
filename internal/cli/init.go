package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdInit(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Clone a remote dotfiles repository",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.CloneRemoteRepo(args[0]); err != nil {
				logger.Error("failed to clone remote repo", "error", err)
			}

			logger.Info("successfully cloned remote repo")
		},
	}
}
