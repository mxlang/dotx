package dotx

import (
	"errors"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
	"path/filepath"
)

type App struct {
	Version string

	AppConfig  *config.AppConfig
	repoConfig *config.RepoConfig
}

func New(version string, appConfig *config.AppConfig, repoConfig *config.RepoConfig) App {
	return App{
		Version:    version,
		AppConfig:  appConfig,
		repoConfig: repoConfig,
	}
}

func (a App) AddDotfile(path string) error {
	filename := filepath.Base(path)
	source := fs.NewPath(path)
	dest := fs.NewPath(filepath.Join(config.RepoDirPath(), filename))

	// check dotfile already added to repo
	if a.repoConfig.DotfileExists(source) {
		return errors.New("dotfile already exists")
	}

	if err := fs.Move(source, dest); err != nil {
		return err
	}

	if err := fs.Symlink(dest, source); err != nil {
		return err
	}

	if err := a.repoConfig.WriteDotfile(source, dest); err != nil {
		return err
	}

	return nil
}

func (a App) DeployDotfiles() error {
	for _, dotfile := range a.repoConfig.Dotfiles {
		source := fs.NewPath(filepath.Join(config.RepoDirPath(), dotfile.Source))
		dest := fs.NewPath(dotfile.Destination)

		if dest.Exists() {
			if dest.IsSymlink() && dest.SymlinkPath() == source.AbsPath() {
				logger.Debug("dotfile already deployed with dotx", "dotfile", source.Filename())
				continue
			} else {
				overwrite := tui.Confirm("File already exists. Overwrite?")
				if !overwrite {
					continue
				}

				if err := fs.Delete(dest); err != nil {
					return err
				}
			}
		}

		// create base dir if not exists
		dir := fs.NewPath(dest.Dir())
		if err := fs.Mkdir(dir); err != nil {
			return err
		}

		if err := fs.Symlink(source, dest); err != nil {
			return err
		}
	}

	return nil
}

func (a App) CloneRemoteRepo(remoteRepo string) error {
	return git.Clone(remoteRepo)
}

func (a App) PullRemoteRepo() error {
	return git.Pull()
}

func (a App) PushRemoteRepo(commitMessage string) error {
	if err := git.Add("."); err != nil {
		return err
	}

	if err := git.Commit(commitMessage); err != nil {
		return err
	}

	if err := git.Push(); err != nil {
		return err
	}

	return nil
}
