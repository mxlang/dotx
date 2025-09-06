package cli

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/spf13/cobra"
)

func newCmdDeploy(app core.App) *cobra.Command {
	var force bool

	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your dotfiles to the current system",
		Long:  "Create symbolic links from your dotfiles to their appropriate locations in your home directory",
		Example: `  dotx deploy
  dotx deploy --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			app.Deploy(force)
		},
	}

	deployCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "never prompt for overwriting")

	return deployCmd
}
