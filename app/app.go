package app

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/fs"
)

const repoDir = ".dotfiles"

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
	return a.fs.Mkdir(a.appConfig.RepoDir)
}

func (a App) AddDotfile(path string) {
	sourcePath, _ := a.fs.AbsPath(path)
	destinationPath := filepath.Join(a.appConfig.RepoDir, filepath.Base(path))

	fmt.Println(sourcePath)
	fmt.Println(destinationPath)

	if err := a.fs.Move(sourcePath, destinationPath); err != nil {
		a.Logger.Fatal(err)
	}

	if err := a.fs.Symlink(destinationPath, sourcePath); err != nil {
		a.Logger.Fatal(err)
	}

	dotfile := config.Dotfile{
		Source:      sourcePath,
		Destination: destinationPath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
		a.Logger.Fatal(err)
	}
}
