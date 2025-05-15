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

	App  AppConfig
	Repo RepoConfig
}

func Load() Config {
	// Ensure the config directory exists
	appDir := fs.NewPath(appDirPath())
	if err := fs.Mkdir(appDir); err != nil {
		logger.Error("error while creating dotx config directory", "error", err)
	}

	// Ensure the repository directory exists
	repoDir := fs.NewPath(repoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating repository directory", "error", err)
	}

	app := loadAppConfig()
	repo := loadRepoConfig()

	return Config{
		RepoPath: repoDirPath(),

		App:  app,
		Repo: repo,
	}
}

func loadAppConfig() AppConfig {
	config := defaultAppConfig()
	path := appConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Warn("invalid config", "error", err)
	}

	return config
}

func defaultAppConfig() AppConfig {
	return AppConfig{
		Verbose:       false,
		CommitMessage: "update dotfiles",
	}
}

func loadRepoConfig() RepoConfig {
	config := RepoConfig{}
	path := repoConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading repository config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Warn("invalid repository config", "error", err)
	}

	return config
}
