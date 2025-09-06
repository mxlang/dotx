package sync

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPush(app core.App) *cobra.Command {
	var commitMessage string

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Save and upload local dotfile changes to remote repository",
		Long:  "Commit local changes to your dotfiles and push them to the remote repository for backup and sharing",
		Example: `  dotx sync push
  dotx sync push -m "Update bash aliases"`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			runPush(app, commitMessage)
		},
	}

	pushCmd.PersistentFlags().StringVarP(&commitMessage, "message", "m", app.Config.CommitMessage, "Specify a commit message")

	return pushCmd
}

func runPush(app core.App, commitMessage string) { // TODO move to core.App or own git struct
	logger.Debug("add changes to dotfiles")
	if err := git.Add(app.Repo.Path, "."); err != nil {
		logger.Error("failed to add changes", "error", err)
	}

	logger.Debug("commit changes to dotfiles", "message", commitMessage)
	if err := git.Commit(app.Repo.Path, commitMessage); err != nil {
		logger.Error("failed to commit changes", "error", err)
	}

	logger.Debug("push changes to dotfiles")
	if err := git.Push(app.Repo.Path); err != nil {
		logger.Error("failed to push changes", "error", err)
	}

	logger.Info("successfully pushed changes to remote dotfiles")
}
