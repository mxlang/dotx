package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
	"github.com/spf13/cobra"
)

func newCmdInit(cfg *config.Config) *cobra.Command {
	var deploy bool
	var force bool

	initCmd := &cobra.Command{
		Use:   "init <repository-url>",
		Short: "Initialize by cloning a remote dotfiles repository",
		Long:  "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files",
		Example: `  dotx sync init https://github.com/username/dotfiles.git
  dotx sync init https://github.com/username/dotfiles.git --deploy`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			runInit(cfg, args[0], deploy, force)
		},
	}

	initCmd.PersistentFlags().BoolVarP(&deploy, "deploy", "d", cfg.App.DeployOnInit, "automatically deploy dotfiles")
	initCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "never prompt for overwriting")

	return initCmd
}

func runInit(cfg *config.Config, url string, deploy bool, force bool) {
	dir := fs.NewPath(cfg.RepoPath)
	if dir.HasSubfiles() {
		overwrite, err := tui.Confirm(
			"Directory is already a Git repository. Overwrite?",
			"",
		)

		if err != nil {
			logger.Error("failed to render TUI", "error", err)
		}

		if !overwrite {
			logger.Debug("overwrite cancelled")
			return
		}

		if err := fs.Delete(dir); err != nil {
			logger.Error("failed to delete", "error", err)
		} else {
			logger.Debug("deleted", "path", dir.AbsPath())
		}
	}

	if err := git.Clone(dir.AbsPath(), url); err != nil {
		logger.Error("failed to clone remote dotfiles", "error", err)
	}

	if deploy {
		logger.Debug("automatically deploy dotfiles was activated")
		runDeploy(cfg, force)
	}

	logger.Info("successfully cloned remote dotfiles")
}
