package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdPull(cfg *config.Config) *cobra.Command {
	var deploy bool
	var force bool

	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Update local dotfiles by pulling changes from remote repository",
		Long:  "Fetch and merge the latest changes from your remote dotfiles repository to keep your local copy up-to-date",
		Example: `  dotx sync pull
  dotx sync pull --deploy
  dotx sync pull --deploy --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			if err := git.Pull(cfg.RepoPath); err != nil {
				logger.Error("failed to pull remote dotfiles", "error", err)
			}

			if deploy {
				logger.Debug("automatically deploy dotfiles was activated")
				runDeploy(cfg, force)
			}

			logger.Info("successfully pulled from remote dotfiles")
		},
	}

	pullCmd.PersistentFlags().BoolVarP(&deploy, "deploy", "d", cfg.App.DeployOnPull, "automatically deploy dotfiles")
	pullCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "never prompt for overwriting")

	return pullCmd
}
