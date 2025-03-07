package config

import (
	"errors"
	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
	"github.com/mlang97/dotx/internal/logger"
	"os"
	"path/filepath"
)

type RepoConfig struct {
	Dotfiles []Dotfile `yaml:"dotfiles"`
}

type Dotfile struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

func LoadRepoConfig() RepoConfig {
	config := RepoConfig{}

	path := filepath.Join(xdg.DataHome, baseDir, repoDir, repoConfigFile)

	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.Debug("config for dotfiles repo not found")
		} else {
			logger.Warn("error while reading dotfiles repo config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Warn("unable to unmarshal dotfiles repo config", "error", err)
	}

	return config
}

func repoDirPath() string {
	return filepath.Join(xdg.ConfigHome, baseDir, repoDir)
}

func repoConfigFilePath() string {
	return filepath.Join(repoDirPath(), repoConfigFile)
}
