package config

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type appConfig struct {
	Verbose     bool
	DotfilesDir string
}

func (a appConfig) CheckDotfilesDir() error {
	fileInfo, err := os.Stat(a.DotfilesDir)
	if err == nil {
		if fileInfo.IsDir() {
			return nil
		}
	}

	return errors.New("dotfiles dir does not exist")
}

func loadAppConfig() appConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	config := appConfig{
		Verbose:     false,
		DotfilesDir: filepath.Join(home, ".dotfiles"),
	}

	appConfigLoader := viper.New()
	appConfigLoader.SetConfigType("yaml")

	// first load ~/.config/dotx/config.yaml
	appConfigLoader.SetConfigName("config")
	appConfigLoader.AddConfigPath(filepath.Join(home, ".config", "dotx"))
	if err := appConfigLoader.ReadInConfig(); err == nil {
		if err := appConfigLoader.Unmarshal(&config); err != nil {
			slog.Warn("failed to unmarshal ~/.config/dotx/config.yaml", "error", err)
		}
		return config
	}

	// fallback to ~/.dotx.yaml
	appConfigLoader.SetConfigName(".dotx")
	appConfigLoader.AddConfigPath(home)
	if err := appConfigLoader.ReadInConfig(); err == nil {
		if err := appConfigLoader.Unmarshal(&config); err != nil {
			slog.Warn("failed to unmarshal ~/.dotx.yaml", "error", err)
		}
		return config
	}

	return config
}
