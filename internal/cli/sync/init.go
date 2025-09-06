package sync

import (
	"slices"

	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/script"
	"github.com/mxlang/dotx/internal/tui"
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
			runInit(app, opts, args[0])
		},
	}

	initCmd.PersistentFlags().BoolVarP(&opts.deploy, "deploy", "d", app.Config.DeployOnInit, "automatically deploy dotfiles")
	initCmd.PersistentFlags().BoolVarP(&opts.force, "force", "f", false, "never prompt for overwriting")

	return initCmd
}

func runInit(app core.App, opts initOptions, url string) { // TODO move to core.App or own git struct
	if shouldCloneDotfiles(app.Repo.Path, url) {
		logger.Debug("clone remote dotfiles", "url", url)
		if err := git.Clone(app.Repo.Path, url); err != nil {
			logger.Error("failed to clone remote dotfiles", "error", err)
		}

		app = app.ReloadRepo()

		logger.Info("successfully cloned remote dotfiles")
	}

	runInitScripts(app)

	if opts.deploy {
		logger.Debug("automatic deploy is active")
		app.Deploy(opts.force)
	}
}

func shouldCloneDotfiles(dir fs.Path, url string) bool {
	remotes, err := git.Remote(dir)
	if err != nil {
		logger.Debug("no remote dotfiles found")
		return true
	}

	if !slices.Contains(remotes, url) {
		overwrite, err := tui.Confirm(
			"Directory is already another Git repository. Overwrite?",
			"",
		)

		if err != nil {
			logger.Error("failed to render TUI", "error", err)
		}

		if !overwrite {
			logger.Debug("overwrite cancelled")
			return false
		}

		logger.Debug("delete", "path", dir)
		if err := fs.Delete(dir); err != nil {
			logger.Error("failed to delete", "error", err)
		}

		return overwrite
	}

	return false
}

func runInitScripts(app core.App) { // TODO refactor see https://github.com/mxlang/dotx/pull/21
	for _, scriptPath := range app.Repo.Scripts.Init {
		fullPath := app.Repo.Path.Join(scriptPath)
		if !fullPath.Exists() {
			logger.Warn("script does not exist", "script", fullPath)
			continue
		}

		logger.Info("execute script", "script", fullPath)
		if err := script.Run(fullPath.AbsPath()); err != nil {
			logger.Warn(err)
		} else {
			logger.Debug("successfully executed script", "script", fullPath)
		}
	}
}
