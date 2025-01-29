package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	RepoDir string
}

func FromAppFile() AppConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	config := AppConfig{
		RepoDir: filepath.Join(home, ".dotfiles"),
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
