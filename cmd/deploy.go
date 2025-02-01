package cmd

import (
	"github.com/mlang97/dotx/app"
	"github.com/spf13/cobra"
)

func newCmdDeploy(dotx app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "A brief description of your command",
		Args:  cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			dotx.DeployDotfiles()
		},
	}
}
