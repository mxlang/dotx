package cmd

import (
	"github.com/mlang97/dotx/app"
	"github.com/spf13/cobra"
)

func newCmdDeploy(dotx app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploys your dotfiles on your system",
		Args:  cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.DeployDotfiles(); err != nil {
				dotx.Logger.Error("failed to deploy your dotfiles", "err", err)
			}

			dotx.Logger.Info("successfully deployed your dotfiles")
		},
	}
}
