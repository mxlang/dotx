package config

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var repoConfigLoader *viper.Viper

type repoConfig struct {
	Dotfiles []Dotfile
}

type Dotfile struct {
	Destination string
	Source      string
}

func (r repoConfig) AddDotfile(destPath string, sourcePath string) {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Error("failed to determine user home directory", "error", err)
		os.Exit(1)
	}

	// sanitize file paths
	destPath = strings.Replace(destPath, home, "$HOME", 1)
	sourcePath = strings.Replace(sourcePath, home+"/.dotfiles/", "", 1)

	file := Dotfile{
		Destination: destPath,
		Source:      sourcePath,
	}
	r.Dotfiles = append(r.Dotfiles, file)

	repoConfigLoader.Set("dotfiles", r.Dotfiles)

	if err := repoConfigLoader.WriteConfig(); err != nil {
		log.Fatal(err)
	}
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
