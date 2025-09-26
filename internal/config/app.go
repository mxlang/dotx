package config

import (
	"errors"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type AppConfig struct {
	Verbose       bool   `yaml:"verbose"`
	CommitMessage string `yaml:"commitMessage"`
	DeployOnInit  bool   `yaml:"deployOnInit"`
	DeployOnPull  bool   `yaml:"deployOnPull"`
}

func LoadAppConfig() AppConfig {
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
			logger.Error("error while reading dotx config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Error("invalid dotx config", "error", err)
	}

	return config
}

func defaultAppConfig() AppConfig {
	return AppConfig{
		Verbose:       false,
		CommitMessage: "update dotfiles",
		DeployOnPull:  false,
		DeployOnInit:  false,
	}
}
