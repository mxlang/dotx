package core

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
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

func (a App) ReloadRepo() App {
	return App{
		Config: a.Config,
		Repo:   config.LoadRepoConfig(),
	}
}

func (a App) Add(dotfile config.Dotfile, optionalDir string) { // TODO should return error
	//if optionalDir != "" {
	//	dir := a.Config.RepoPath.Join(optionalDir)
	//	if err := fs.Mkdir(dir); err != nil {
	//		logger.Error("could not create directory", "dir", dir, "error", err)
	//	}
	//	source = cfg.RepoPath.Join(optionalDir, filename)
	//}

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
}
