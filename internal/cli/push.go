package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPush(dotx dotx.App) *cobra.Command {
	var commitMessage string

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Save and upload local dotfile changes to remote repository",
		Long:  "Commit local changes to your dotfiles and push them to the remote repository for backup and sharing",
		Example: `  dotx sync push
  dotx sync push -m "Update bash aliases"`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.PushRemoteRepo(commitMessage); err != nil {
				logger.Error("failed to push to remote repo", "error", err)
			}

			logger.Info("successfully pushed to remote repo")
		},
	}

	pushCmd.PersistentFlags().StringVarP(&commitMessage, "message", "m", dotx.AppConfig.CommitMessage, "Specify a commit message")

	return pushCmd
}
