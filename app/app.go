package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/fs"
)

type App struct {
	Logger *log.Logger
	fs     fs.Filesystem

	appConfig  config.AppConfig
	repoConfig config.RepoConfig
}

func New(logger *log.Logger, fs fs.Filesystem, appConfig config.AppConfig, repoConfig config.RepoConfig) App {
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

func (a App) AddDotfile(path string) {
	fileName := filepath.Base(path)

	if a.repoConfig.GetDotfile(fileName) != (config.Dotfile{}) {
		a.Logger.Fatal("dotfile already exist")
	}

	sourcePath, _ := a.fs.AbsPath(path)
	destinationPath := filepath.Join(a.appConfig.GetRepoDir(), fileName)

	if err := a.fs.Move(sourcePath, destinationPath); err != nil {
		a.Logger.Fatal(err)
	}

	if err := a.fs.Symlink(destinationPath, sourcePath); err != nil {
		a.Logger.Fatal(err)
	}

	// normalize paths
	home, _ := os.UserHomeDir()
	sourcePath = strings.Replace(sourcePath, home, "$HOME", 1)

	dotfile := config.Dotfile{
		Source:      fileName,
		Destination: sourcePath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
		a.Logger.Fatal(err)
	}
}
