package config

import (
	"errors"
	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"os"
	"path/filepath"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
)

func appDirPath() string {
	return filepath.Join(xdg.ConfigHome, baseDir)
}

func appConfigFilePath() string {
	return filepath.Join(appDirPath(), appConfigFile)
}

func RepoDirPath() string {
	return filepath.Join(xdg.DataHome, baseDir, repoDir)
}

func repoConfigFilePath() string {
	return filepath.Join(RepoDirPath(), repoConfigFile)
}

type Config struct {
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
	repoDir := fs.NewPath(RepoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating repository directory", "error", err)
	}

	app := loadAppConfig()
	repo := loadRepoConfig()

	return Config{
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

	if err := yaml.Unmarshal(content, config); err != nil {
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

	if err := yaml.Unmarshal(content, config); err != nil {
		logger.Warn("invalid repository config", "error", err)
	}

	return config
}
