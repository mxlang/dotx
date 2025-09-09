package sync

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/spf13/cobra"
)

type pullOptions struct {
	deploy bool
	force  bool
}

func newCmdPull(app core.App) *cobra.Command {
	opts := pullOptions{}

	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Update local dotfiles by pulling changes from remote repository",
		Long:  "Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date",
		Example: `  dotx sync pull
  dotx sync pull --deploy
  dotx sync pull --deploy --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			app.Pull(opts.deploy, opts.force)
		},
	}

	pullCmd.PersistentFlags().BoolVarP(&opts.deploy, "deploy", "d", app.Config.DeployOnPull, "automatically deploy dotfiles")
	pullCmd.PersistentFlags().BoolVarP(&opts.force, "force", "f", false, "never prompt for overwriting")

	return pullCmd
}
