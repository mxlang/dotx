package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPush(cfg config.Config) *cobra.Command {
	var commitMessage string

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Save and upload local dotfile changes to remote repository",
		Long:  "Commit local changes to your dotfiles and push them to the remote repository for backup and sharing",
		Example: `  dotx sync push
  dotx sync push -m "Update bash aliases"`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := git.Add(cfg.RepoPath, "."); err != nil {
				logger.Error("failed to add changes: %w", err)
			}

			if err := git.Commit(cfg.RepoPath, commitMessage); err != nil {
				logger.Error("failed to commit changes: %w", err)
			}

			if err := git.Push(cfg.RepoPath); err != nil {
				logger.Error("failed to push changes: %w", err)
			}

			logger.Info("successfully pushed to remote repository")
		},
	}

	pushCmd.PersistentFlags().StringVarP(&commitMessage, "message", "m", cfg.App.CommitMessage, "Specify a commit message")

	return pushCmd
}
