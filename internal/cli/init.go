package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/script"
	"github.com/mxlang/dotx/internal/tui"
	"github.com/spf13/cobra"
	"path/filepath"
)

type initOptions struct {
	deploy bool
	force  bool
}

func newCmdInit(cfg *config.Config) *cobra.Command {
	opts := initOptions{}

	initCmd := &cobra.Command{
		Use:   "init <repository-url>",
		Short: "Initialize by cloning a remote dotfiles repository",
		Long:  "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files",
		Example: `  dotx sync init https://github.com/username/dotfiles.git
  dotx sync init https://github.com/username/dotfiles.git --deploy
  dotx sync init https://github.com/username/dotfiles.git --deploy --force`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			runInit(cfg, opts, args[0])
		},
	}

	initCmd.PersistentFlags().BoolVarP(&opts.deploy, "deploy", "d", cfg.App.DeployOnInit, "automatically deploy dotfiles")
	initCmd.PersistentFlags().BoolVarP(&opts.force, "force", "f", false, "never prompt for overwriting")

	return initCmd
}

func runInit(cfg *config.Config, opts initOptions, url string) {
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

		logger.Debug("delete", "path", dir.AbsPath())
		if err := fs.Delete(dir); err != nil {
			logger.Error("failed to delete", "error", err)
		}
	}

	logger.Debug("clone remote dotfiles", "url", url)
	if err := git.Clone(dir.AbsPath(), url); err != nil {
		logger.Error("failed to clone remote dotfiles", "error", err)
	}

	logger.Info("successfully cloned remote dotfiles")

	runInitScripts(cfg)

	if opts.deploy {
		logger.Debug("automatic deploy is active")
		runDeploy(cfg, opts.force)
	}
}

func runInitScripts(cfg *config.Config) {
	for _, scriptPath := range cfg.Repo.Scripts.Init {
		fullPath := fs.NewPath(filepath.Join(cfg.RepoPath, scriptPath))
		if !fullPath.Exists() {
			logger.Warn("script does not exist", "script", fullPath.AbsPath())
			continue
		}

		logger.Info("execute script", "script", fullPath.AbsPath())
		if err := script.Run(fullPath.AbsPath()); err != nil {
			logger.Warn(err)
		} else {
			logger.Debug("successfully executed script", "script", fullPath.AbsPath())
		}
	}
}
