package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPush(cfg *config.Config) *cobra.Command {
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
				logger.Error("failed to add changes", "error", err)
			} else {
				logger.Debug("successfully added changes to dotfiles")
			}

			if err := git.Commit(cfg.RepoPath, commitMessage); err != nil {
				logger.Error("failed to commit changes", "error", err)
			} else {
				logger.Debug("successfully committed changes to dotfiles", "message", commitMessage)
			}

			if err := git.Push(cfg.RepoPath); err != nil {
				logger.Error("failed to push changes", "error", err)
			}

			logger.Info("successfully pushed changes to remote dotfiles")
		},
	}

	pushCmd.PersistentFlags().StringVarP(&commitMessage, "message", "m", cfg.App.CommitMessage, "Specify a commit message")

	return pushCmd
}
