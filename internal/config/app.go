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

type AppConfig struct {
	Verbose       bool   `yaml:"verbose"`
	CommitMessage string `yaml:"commitMessage"`
}

func LoadAppConfig() *AppConfig {
	ensureAppConfigDir()

	config := defaultAppConfig()
	path := appConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// TODO debug is not working because of different log level at startup
			logger.Debug("config for dotx not found")
		} else {
			logger.Warn("error while reading dotx config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, config); err != nil {
		logger.Warn("unable to unmarshal dotx config", "error", err)
	}

	return config
}

func ensureAppConfigDir() {
	appDir := fs.NewPath(appDirPath())
	if err := fs.Mkdir(appDir); err != nil {
		logger.Error("error while creating dotx config dir", "error", err)
	} else {
		logger.Debug("dotx config dir created or already existent", "dir", appDir.AbsPath())
	}
}

func defaultAppConfig() *AppConfig {
	return &AppConfig{
		Verbose:       false,
		CommitMessage: "update dotfiles",
	}
}

func appDirPath() string {
	return filepath.Join(xdg.ConfigHome, baseDir)
}

func appConfigFilePath() string {
	return filepath.Join(appDirPath(), appConfigFile)
}
