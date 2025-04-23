package config

import (
	"errors"
	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"os"
	"path/filepath"
	"strings"
)

type RepoConfig struct {
	Dotfiles []Dotfile `yaml:"dotfiles"`
}

type Dotfile struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

func (r *RepoConfig) DotfileExists(source fs.Path) bool {
	for _, dotfile := range r.Dotfiles {
		if fs.NewPath(dotfile.Destination) == source {
			return true
		}
	}

	return false
}

func (r *RepoConfig) WriteDotfile(source, dest fs.Path) error {
	// normalize paths
	home, _ := os.UserHomeDir()
	sourcePath := strings.Replace(source.AbsPath(), home, "$HOME", 1)
	destinationPath := strings.Replace(dest.AbsPath(), RepoDirPath(), "", 1)

	dotfile := Dotfile{
		Source:      destinationPath,
		Destination: sourcePath,
	}

	r.Dotfiles = append(r.Dotfiles, dotfile)

	config, err := yaml.Marshal(r)
	if err != nil {
		return errors.New("unable to marshal dotfiles repo config")
	}

	if err := os.WriteFile(repoConfigFilePath(), config, 0644); err != nil {
		return errors.New("unable to write dotfiles repo config")
	}

	return nil
}

func LoadRepoConfig() *RepoConfig {
	ensureRepoConfigDir()

	config := &RepoConfig{}
	path := repoConfigFilePath()

	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			logger.Debug("config for dotfiles repo not found")
		} else {
			logger.Warn("error while reading dotfiles repo config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, config); err != nil {
		logger.Warn("unable to unmarshal dotfiles repo config", "error", err)
	}

	return config
}

func ensureRepoConfigDir() {
	repoDir := fs.NewPath(RepoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating dotfiles repo dir", "error", err)
	} else {
		logger.Debug("dotfiles repo dir created or already existent", "dir", repoDir.AbsPath())
	}
}

func RepoDirPath() string {
	return filepath.Join(xdg.DataHome, baseDir, repoDir)
}

func repoConfigFilePath() string {
	return filepath.Join(RepoDirPath(), repoConfigFile)
}
