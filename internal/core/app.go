package core

import (
	"slices"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
)

type App struct {
	Config config.AppConfig
	Repo   config.RepoConfig
}

func NewApp(config config.AppConfig, repo config.RepoConfig) App {
	return App{
		Config: config,
		Repo:   repo,
	}
}

func (a App) Add(dotfile config.Dotfile, optionalDir string) { // TODO should return error
	if optionalDir != "" {
		dir := a.Repo.Path.Join(optionalDir)
		if err := fs.Mkdir(dir); err != nil {
			logger.Error("could not create directory", "dir", dir, "error", err)
		}
		dotfile.Source = a.Repo.Path.Join(optionalDir, dotfile.Source.Filename())
	}

	if a.Repo.HasDotfile(dotfile) {
		logger.Error("already exists in dotfiles", "dotfile", dotfile.Source.Filename())
	}

	logger.Debug("move", "from", dotfile.Destination, "to", dotfile.Source)
	if err := fs.Move(dotfile.Destination, dotfile.Source); err != nil {
		logger.Error("failed to move", "error", err)
	}

	logger.Debug("create symlink", "from", dotfile.Source, "to", dotfile.Destination)
	if err := fs.Symlink(dotfile.Source, dotfile.Destination); err != nil {
		logger.Error("failed to create symlink", "error", err)
	}

	logger.Debug("write to dotfiles config")
	if err := a.Repo.AddDotfile(dotfile); err != nil {
		logger.Error("failed to write dotfiles config", "error", err)
	}

	logger.Info("successfully added", "dotfile", dotfile.Source.Filename())
}

func (a App) Deploy(force bool) { // TODO should return error
	for _, dotfile := range a.Repo.Dotfiles {
		logger.Debug("deploy dotfile", "from", dotfile.Source, "to", dotfile.Destination)

		if dotfile.Destination.Exists() {
			if dotfile.Destination.IsSymlink() && dotfile.Destination.SymlinkPath() == dotfile.Source.AbsPath() {
				logger.Debug("dotfile already deployed with dotx", "dotfile", dotfile.Source.Filename())
				continue
			}

			if !force {
				title := "File already exists. Overwrite?"
				if dotfile.Destination.IsDir() {
					title = "Directory already exists. Overwrite?"
				}

				overwrite, err := tui.Confirm(
					title,
					dotfile.Destination.AbsPath(),
				)

				if err != nil {
					logger.Error("failed to render TUI", "error", err)
				}

				if !overwrite {
					logger.Debug("overwrite cancelled")
					continue
				}
			}

			logger.Debug("delete", "path", dotfile.Destination)
			if err := fs.Delete(dotfile.Destination); err != nil {
				logger.Error("failed to delete", "error", err)
			}
		}

		// Ensure parent directory exists
		dir := fs.NewPath(dotfile.Destination.Dir())
		logger.Debug("create parent directory if not exists", "dir", dir)
		if err := fs.Mkdir(dir); err != nil {
			logger.Error("could not create parent directory", "error", err)
		}

		logger.Debug("create symlink", "from", dotfile.Source, "to", dotfile.Destination)
		if err := fs.Symlink(dotfile.Source, dotfile.Destination); err != nil {
			logger.Error("failed to create symlink", "error", err)
		}

		logger.Info("successfully deployed", "dotfile", dotfile.Source.Filename())
	}

	a.Repo.ExecuteScripts(config.OnDeploy)
}

func (a App) Init(url string, deploy bool, force bool) { // TODO should return error
	if shouldCloneDotfiles(a.Repo.Path, url) {
		logger.Debug("clone remote dotfiles", "url", url)
		if err := git.Clone(a.Repo.Path, url); err != nil {
			logger.Error("failed to clone remote dotfiles", "error", err)
		}

		a.Repo = config.LoadRepoConfig()

		logger.Info("successfully cloned remote dotfiles")
	}

	a.Repo.ExecuteScripts(config.OnInit)

	if deploy {
		logger.Debug("automatic deploy is active")
		a.Deploy(force)
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

func (a App) Pull(deploy bool, force bool) { // TODO should return error
	logger.Debug("pull changes from remote dotfiles")
	if err := git.Pull(a.Repo.Path); err != nil {
		logger.Error("failed to pull remote dotfiles", "error", err)
	}

	logger.Info("successfully pulled from remote dotfiles")

	a.Repo.ExecuteScripts(config.OnPull)

	if deploy {
		logger.Debug("automatic deploy is active")
		a.Deploy(force)
	}
}

func (a App) Push(commitMessage string) { // TODO should return error
	logger.Debug("add changes to dotfiles")
	if err := git.Add(a.Repo.Path, "."); err != nil {
		logger.Error("failed to add changes", "error", err)
	}

	logger.Debug("commit changes to dotfiles", "message", commitMessage)
	if err := git.Commit(a.Repo.Path, commitMessage); err != nil {
		logger.Error("failed to commit changes", "error", err)
	}

	logger.Debug("push changes to dotfiles")
	if err := git.Push(a.Repo.Path); err != nil {
		logger.Error("failed to push changes", "error", err)
	}

	logger.Info("successfully pushed changes to remote dotfiles")
}
