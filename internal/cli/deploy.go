package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdDeploy(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploys your dotfiles on your system",
		Args:  cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.DeployDotfiles(); err != nil {
				logger.Error("failed to deploy your dotfiles", "error", err)
			}

			logger.Info("successfully deployed your dotfiles")
		},
	}
}
