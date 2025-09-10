package config

import (
	"errors"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type appConfig struct {
	Verbose       bool   `yaml:"verbose"`
	CommitMessage string `yaml:"commitMessage"`
	DeployOnInit  bool   `yaml:"deployOnInit"`
	DeployOnPull  bool   `yaml:"deployOnPull"`
}

func defaultAppConfig() appConfig {
	return appConfig{
		Verbose:       false,
		CommitMessage: "update dotfiles",
		DeployOnPull:  false,
		DeployOnInit:  false,
	}
}

func loadAppConfig() appConfig {
	// Ensure the config directory exists
	appDir := fs.NewPath(appDirPath())
	if err := fs.Mkdir(appDir); err != nil {
		logger.Error("error while creating dotx config directory", "error", err)
	}

	config := defaultAppConfig()
	content, err := os.ReadFile(appConfigFilePath())
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
