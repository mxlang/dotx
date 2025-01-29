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

func (a AppConfig) GetRepoDir() string {
	return os.ExpandEnv(a.RepoDir)
}

func FromAppFile() AppConfig {
	home, _ := os.UserHomeDir()
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
