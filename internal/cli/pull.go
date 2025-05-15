package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPull(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Short:   "Update local dotfiles by pulling changes from remote repository",
		Long:    "Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date",
		Example: `  dotx sync pull`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := git.Pull(cfg.RepoPath); err != nil {
				logger.Error("failed to pull remote repository: %w", err)
			}

			logger.Info("successfully pulled from remote repository")
		},
	}
}
