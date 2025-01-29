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
	sourcePath, _ := a.fs.AbsPath(path)
	destinationPath := filepath.Join(a.appConfig.GetRepoDir(), filepath.Base(path))

	if err := a.fs.Move(sourcePath, destinationPath); err != nil {
		a.Logger.Fatal(err)
	}

	if err := a.fs.Symlink(destinationPath, sourcePath); err != nil {
		a.Logger.Fatal(err)
	}

	// normalize paths
	home, _ := os.UserHomeDir()
	sourcePath = strings.Replace(sourcePath, home, "$HOME", 1)
	destinationPath = strings.Replace(destinationPath, a.appConfig.GetRepoDir()+"/", "", 1)

	dotfile := config.Dotfile{
		Source:      sourcePath,
		Destination: destinationPath,
	}

	if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
		a.Logger.Fatal(err)
	}
}
