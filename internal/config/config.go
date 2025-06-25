package config

import (
	"errors"
	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"os"
)

type Config struct {
	RepoPath string

	App  appConfig
	Repo repoConfig
}

func Load() *Config {
	app := loadAppConfig()
	repo := loadRepoConfig()

	return &Config{
		RepoPath: repoDirPath(),

		App:  app,
		Repo: repo,
	}
}

func loadAppConfig() appConfig {
	// Ensure the config directory exists
	appDir := fs.NewPath(appDirPath())
	if err := fs.Mkdir(appDir); err != nil {
		logger.Error("error while creating dotx config directory", "error", err)
	}

	config := defaultAppConfig()
	path := appConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading dotx config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Warn("invalid dotx config", "error", err)
	}

	return config
}

func defaultAppConfig() appConfig {
	return appConfig{
		Verbose:       false,
		CommitMessage: "update dotfiles",
		DeployOnPull:  false,
		DeployOnInit:  false,
	}
}

func loadRepoConfig() repoConfig {
	// Ensure the dotfiles directory exists
	repoDir := fs.NewPath(repoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating dotfiles directory", "error", err)
	}

	config := repoConfig{}
	path := repoConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading dotfiles config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Warn("invalid dotfiles config", "error", err)
	}

	return config
}
