package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
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
		return errors.New("failed to write config")
	}

	return nil
}

func (a App) DeployDotfiles() error {
	for _, dotfile := range a.repoConfig.Dotfiles {
		sourcePath := filepath.Join(a.appConfig.RepoDir, dotfile.Source)
		destPath := os.ExpandEnv(dotfile.Destination)

		// TODO check if file is already deployed from dotx

		_, err := os.Stat(destPath)
		if err == nil {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Dotfile already exists on your sytem. Do you want to backup and overwrite? (Y/n): ")
			char, _, err := reader.ReadRune()
			if err != nil {
				return err
			}

			switch char {
			case 'y', 'Y', '\n':
			// TODO move existing file to backup folder
			case 'n', 'N':
				continue
			default:
				return errors.New("invalid user input")
			}
		}

		if err := a.fs.Symlink(sourcePath, destPath); err != nil {
			return errors.New("failed to symlink file")
		}
	}

	return nil
}

func (a App) InitializeRemoteRepo(remoteRepo string) error {
	command := exec.Command("git", "clone", remoteRepo, a.appConfig.RepoDir)
	if err := command.Run(); err != nil {
		return errors.New("failed to clone remote repo")
	}

	return nil
}
