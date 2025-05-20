package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

type pullOptions struct {
	deploy bool
	force  bool
}

func newCmdPull(cfg *config.Config) *cobra.Command {
	pullOpts := pullOptions{}

	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Update local dotfiles by pulling changes from remote repository",
		Long:  "Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date",
		Example: `  dotx sync pull
  dotx sync pull --deploy
  dotx sync pull --deploy --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			runPull(cfg, pullOpts)
		},
	}

	pullCmd.PersistentFlags().BoolVarP(&pullOpts.deploy, "deploy", "d", cfg.App.DeployOnPull, "automatically deploy dotfiles")
	pullCmd.PersistentFlags().BoolVarP(&pullOpts.force, "force", "f", false, "never prompt for overwriting")

	return pullCmd
}

func runPull(cfg *config.Config, pullOpts pullOptions) {
	if err := git.Pull(cfg.RepoPath); err != nil {
		logger.Error("failed to pull remote dotfiles", "error", err)
	}

	if pullOpts.deploy {
		logger.Debug("automatically deploy dotfiles was activated")
		runDeploy(cfg, pullOpts.force)
	}

	logger.Info("successfully pulled from remote dotfiles")
}
