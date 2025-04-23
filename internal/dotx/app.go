package dotx

import (
	"errors"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"path/filepath"
)

type App struct {
	AppConfig  *config.AppConfig
	repoConfig *config.RepoConfig
}

func New(appConfig *config.AppConfig, repoConfig *config.RepoConfig) App {
	return App{
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
				// TODO log warning instead of return error
				return errors.New("dotfile already deployed with dotx")
			} else {
				// TODO ask for overwrite with a TUI
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
