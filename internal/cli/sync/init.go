package sync

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

type initOptions struct {
	deploy bool
	force  bool
}

func newCmdInit(app core.App) *cobra.Command {
	opts := initOptions{}

	initCmd := &cobra.Command{
		Use:   "init [repository-url]",
		Short: "Initialize by cloning a remote dotfiles repository",
		Long:  "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files and running your configured scripts",
		Example: `  dotx sync init https://github.com/username/dotfiles.git
  dotx sync init https://github.com/username/dotfiles.git --deploy
  dotx sync init https://github.com/username/dotfiles.git --deploy --force`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Init(args[0], opts.deploy, opts.force); err != nil {
				logger.Error("failed to initialize dotfiles", "error", err)
			}
		},
	}

	initCmd.PersistentFlags().BoolVarP(&opts.deploy, "deploy", "d", app.Config.DeployOnInit, "automatically deploy dotfiles")
	initCmd.PersistentFlags().BoolVarP(&opts.force, "force", "f", false, "never prompt for overwriting")

	return initCmd
}
