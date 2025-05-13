package dotx

import (
	"errors"
	"fmt"
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
		return fmt.Errorf("failed to move: %w", err)
	}

	if err := fs.Symlink(dest, source); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	if err := a.repoConfig.WriteDotfile(source, dest); err != nil {
		return fmt.Errorf("failed to write repository config: %w", err)
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
			}

			title := "File already exists. Overwrite?"
			if dest.IsDir() {
				title = "Directory already exists. Overwrite?"
			}

			overwrite := tui.Confirm(
				title,
				dest.AbsPath(),
			)
			if !overwrite {
				continue
			}

			if err := fs.Delete(dest); err != nil {
				return fmt.Errorf("failed to delete: %w", err)
			}
		}

		// Ensure parent directory exists
		dir := fs.NewPath(dest.Dir())
		if err := fs.Mkdir(dir); err != nil {
			return fmt.Errorf("could not create parent directory for %s: %w", dest.AbsPath(), err)
		}

		if err := fs.Symlink(source, dest); err != nil {
			return fmt.Errorf("failed to create symlink: %w", err)
		}
	}

	return nil
}

func (a App) CloneRemoteRepo(remoteRepo string) error {
	if err := git.Clone(remoteRepo); err != nil {
		return fmt.Errorf("failed to clone remote repository: %w", err)
	}

	return nil
}

func (a App) PullRemoteRepo() error {
	if err := git.Pull(); err != nil {
		return fmt.Errorf("failed to pull remote repository: %w", err)
	}

	return nil
}

func (a App) PushRemoteRepo(commitMessage string) error {
	if err := git.Add("."); err != nil {
		return fmt.Errorf("failed to add changes: %w", err)
	}

	if err := git.Commit(commitMessage); err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	if err := git.Push(); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	return nil
}
