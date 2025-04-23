package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPull(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Short:   "Update local dotfiles by pulling changes from remote repository",
		Long:    "Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date",
		Args:    cobra.NoArgs,
		Example: "  dotx sync pull",

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.PullRemoteRepo(); err != nil {
				logger.Error("failed to pull from remote repo", "error", err)
			}

			logger.Info("successfully pulled from remote repo")
		},
	}
}
