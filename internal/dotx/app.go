package dotx

import (
	"errors"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/fs"
	"os"
	"path/filepath"
	"strings"
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

	// normalize paths
	// TODO refactor move to WriteDotfile
	home, _ := os.UserHomeDir()
	sourcePath := strings.Replace(source.AbsPath(), home, "$HOME", 1)
	destinationPath := strings.Replace(dest.AbsPath(), config.RepoDirPath(), "", 1)

	dotfile := config.Dotfile{
		Source:      destinationPath,
		Destination: sourcePath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
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
			return errors.New("failed to create dir")
		}

		if err := fs.Symlink(source, dest); err != nil {
			return errors.New("failed to symlink file")
		}
	}

	return nil
}

func (a App) InitializeRemoteRepo(remoteRepo string) error {
	//command := exec.Command("git", "clone", remoteRepo, a.appConfig.RepoDir)
	//if err := command.Run(); err != nil {
	//	return errors.New("failed to clone remote repo")
	//}

	return nil
}
