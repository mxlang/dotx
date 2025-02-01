package app

import (
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
	return a.fs.Mkdir(a.appConfig.GetRepoDir())
}

func (a App) AddDotfile(path string, optDir string) {
	fileName := filepath.Base(path)
	if a.repoConfig.GetDotfile(fileName) != (config.Dotfile{}) {
		a.Logger.Error("dotfile already exist")
	}

	sourcePath, _ := a.fs.AbsPath(path)
	destinationPath := filepath.Join(a.appConfig.GetRepoDir(), fileName)

	if optDir != "" {
		dir := filepath.Join(a.appConfig.GetRepoDir(), optDir)
		if err := a.fs.Mkdir(dir); err != nil {
			a.Logger.Error("failed to create dir")
		}

		destinationPath = filepath.Join(dir, fileName)
	}

	if err := a.fs.Move(sourcePath, destinationPath); err != nil {
		a.Logger.Error(err)
	}

	if err := a.fs.Symlink(destinationPath, sourcePath); err != nil {
		a.Logger.Error(err)
	}

	// normalize paths
	home, _ := os.UserHomeDir()
	sourcePath = strings.Replace(sourcePath, home, "$HOME", 1)
	destinationPath = strings.Replace(destinationPath, a.appConfig.GetRepoDir(), "", 1)

	dotfile := config.Dotfile{
		Source:      destinationPath,
		Destination: sourcePath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
		a.Logger.Error(err)
	}
}

func (a App) DeployDotfiles() {
	for _, dotfile := range a.repoConfig.Dotfiles {
		fmt.Println(dotfile)
		source := filepath.Join(a.appConfig.GetRepoDir(), dotfile.Source)
		dest := os.ExpandEnv(dotfile.Destination)

		a.fs.Symlink(source, dest)
	}
}
