package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPull(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:   "pull",
		Short: "Pull changes from your remote dotfiles repository",
		Args:  cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.PullRemoteRepo(); err != nil {
				logger.Error("failed to pull from remote repo", "error", err)
			}

			logger.Info("successfully pulled from remote repo")
		},
	}
}
