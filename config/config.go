package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	viper.SetConfigType("yaml")

	// first load ~/.config/dotx/config.yaml
	configPath := filepath.Join(home, ".config", "dotx")
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err == nil {
		slog.Debug("loaded config file", "config", filepath.Join(configPath, "config.yaml"))
		return
	}
	slog.Debug("config file not found", "config", filepath.Join(configPath, "config.yaml"))

	// fallback to ~/.dotx.yaml
	viper.SetConfigName(".dotx")
	viper.AddConfigPath(home)
	if err := viper.ReadInConfig(); err == nil {
		slog.Debug("loaded fallback config file", "config", filepath.Join(home, ".dotx.yaml"))
		return
	}
	slog.Debug("fallback config file not found", "config", filepath.Join(home, ".dotx.yaml"))
}
