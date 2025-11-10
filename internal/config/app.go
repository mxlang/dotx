package config

import (
	"errors"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type AppConfig struct {
	Verbose      bool `yaml:"verbose"`
	DeployOnInit bool `yaml:"deployOnInit"`
	DeployOnPull bool `yaml:"deployOnPull"`
}

func LoadAppConfig() AppConfig {
	// Ensure the config directory exists
	if err := fs.Mkdir(appDirPath()); err != nil {
		logger.Error("error while creating dotx config directory", "error", err)
	}

	config := defaultAppConfig()
	path := appConfigFilePath()

	content, err := os.ReadFile(path.AbsPath())
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
		Verbose:      false,
		DeployOnPull: false,
		DeployOnInit: false,
	}
}
