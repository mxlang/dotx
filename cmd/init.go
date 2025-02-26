package cmd

import (
	"github.com/mlang97/dotx/internal/app"
	"github.com/spf13/cobra"
)

func newCmdInit(dotx app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Command to handle everything with an remote repo",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.InitializeRemoteRepo(args[0]); err != nil {
				dotx.Logger.Error("failed to init remote repo", "error", err)
			}

			// TODO create .gitignore if not exists and add backup folder to it

			dotx.Logger.Info("successfully init remote repo")
		},
	}
}
