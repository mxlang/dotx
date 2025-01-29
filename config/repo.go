package config

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/mlang97/dotx/dotfile"
	"github.com/spf13/viper"
)

var repoConfigLoader *viper.Viper

type repoConfig struct {
	Dotfiles []dotfile.Dotfile
}

func (r repoConfig) AddDotfile(dotfile dotfile.Dotfile) error {
	if slices.Contains(r.Dotfiles, dotfile) {
		return errors.New("dotfile already added")
	}

	repoConfigLoader.Set("dotfiles", append(r.Dotfiles, dotfile))

	if err := repoConfigLoader.WriteConfig(); err != nil {
		return errors.New("failed to write config")
	}

	return nil
}

func (r repoConfig) EnsureRepoConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	repoConfigPath := filepath.Join(home, ".dotfiles", "dotx.yaml")

	if _, err := os.Stat(repoConfigPath); err != nil {
		if _, err := os.Create(repoConfigPath); err != nil {
			log.Fatal(err)
		}
	}
}

func loadRepoConfig() repoConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	config := repoConfig{}

	repoConfigLoader.SetConfigName("dotx")
	repoConfigLoader.SetConfigType("yaml")

	repoConfigLoader.AddConfigPath(filepath.Join(home, ".dotfiles"))
	if err := repoConfigLoader.ReadInConfig(); err == nil {
		if err := repoConfigLoader.Unmarshal(&config); err != nil {
			slog.Warn("failed to unmarshal ~/.dotfiles/dotx.yaml", "error", err)
		}
		return config
	}

	return config
}

func init() {
	repoConfigLoader = viper.New()
}
