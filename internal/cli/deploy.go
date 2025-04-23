package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdDeploy(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy your dotfiles to the current system",
		Long:    "Create symbolic links from your dotfiles repository to their appropriate locations in your home directory",
		Args:    cobra.NoArgs,
		Example: "  dotx deploy",

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.DeployDotfiles(); err != nil {
				logger.Error("failed to deploy your dotfiles", "error", err)
			}

			logger.Info("successfully deployed your dotfiles")
		},
	}
}
