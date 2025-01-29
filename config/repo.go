package config

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var repoConfigLoader *viper.Viper

type repoConfig struct {
	appConfig appConfig

	Dotfiles []Dotfile
}

func (r repoConfig) AddDotfile(path string) error {
	dotfile, err := New(r.appConfig, path)
	if err != nil {
		return err
	}

	if err := dotfile.move(); err != nil {
		return err
	}

	if err := dotfile.symlink(); err != nil {
		return err
	}

	r.Dotfiles = append(r.Dotfiles, dotfile)

	repoConfigLoader.Set("dotfiles", r.Dotfiles)

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

func loadRepoConfig(appConfig appConfig) repoConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	config := repoConfig{
		appConfig: appConfig,
	}

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
