package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/fs"
	"github.com/mlang97/dotx/log"
)

type App struct {
	Logger log.Logger
	fs     fs.Filesystem

	appConfig  config.AppConfig
	repoConfig config.RepoConfig
}

func New(logger log.Logger, fs fs.Filesystem, appConfig config.AppConfig, repoConfig config.RepoConfig) App {
	return App{
		Logger:     logger,
		fs:         fs,
		appConfig:  appConfig,
		repoConfig: repoConfig,
	}
}

func (a App) EnsureRepo() error {
	return a.fs.Mkdir(a.appConfig.RepoDir)
}

func (a App) AddDotfile(path string, optDir string) error {
	fileName := filepath.Base(path)
	sourcePath, _ := a.fs.AbsPath(path)
	destinationPath := filepath.Join(a.appConfig.RepoDir, fileName)

	if a.repoConfig.GetDotfile(sourcePath) != (config.Dotfile{}) {
		return errors.New("dotfile already exist")
	}

	if optDir != "" {
		dir := filepath.Join(a.appConfig.RepoDir, optDir)
		if err := a.fs.Mkdir(dir); err != nil {
			return errors.New("failed to create dir")
		}

		destinationPath = filepath.Join(dir, fileName)
	}

	if err := a.fs.Move(sourcePath, destinationPath); err != nil {
		return errors.New("failed to move file")
	}

	if err := a.fs.Symlink(destinationPath, sourcePath); err != nil {
		return errors.New("failed to symlink file")
	}

	// normalize paths
	home, _ := os.UserHomeDir()
	sourcePath = strings.Replace(sourcePath, home, "$HOME", 1)
	destinationPath = strings.Replace(destinationPath, a.appConfig.RepoDir, "", 1)

	dotfile := config.Dotfile{
		Source:      destinationPath,
		Destination: sourcePath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
		return errors.New("failed to write confi")
	}

	return nil
}

func (a App) DeployDotfiles() {
	for _, dotfile := range a.repoConfig.Dotfiles {
		fmt.Println(dotfile)
		source := filepath.Join(a.appConfig.RepoDir, dotfile.Source)
		dest := os.ExpandEnv(dotfile.Destination)

		a.fs.Symlink(source, dest)
	}
}
