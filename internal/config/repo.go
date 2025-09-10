package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type dotfile struct {
	Source      string `yaml:"source"`      // TODO change type to fs.Path
	Destination string `yaml:"destination"` // TODO change type to fs.Path
}

type repoConfig struct {
	Dotfiles []dotfile `yaml:"dotfiles"`
	Scripts  []script  `yaml:"scripts"`
}

func (r *repoConfig) HasDotfile(source fs.Path) bool {
	for _, dotfile := range r.Dotfiles {
		if fs.NewPath(dotfile.Destination) == source {
			return true
		}
	}

	return false
}

func (r *repoConfig) AddDotfile(source fs.Path, dest fs.Path) error {
	// Normalize paths
	home, _ := os.UserHomeDir()
	sourcePath := strings.Replace(source.AbsPath(), home, "$HOME", 1)
	destinationPath := strings.Replace(dest.AbsPath(), repoDirPath(), "", 1)

	dotfile := dotfile{
		Source:      destinationPath,
		Destination: sourcePath,
	}

	r.Dotfiles = append(r.Dotfiles, dotfile)

	config, err := yaml.Marshal(r)
	if err != nil {
		return fmt.Errorf("unable to marshal dotfiles config: %w", err)
	}

	if err := os.WriteFile(repoConfigFilePath(), config, 0644); err != nil {
		return fmt.Errorf("unable to write dotfiles config: %w", err)
	}

	return nil
}

func (r *repoConfig) ExecuteScripts(event event) {
	for _, script := range r.Scripts {
		script.execute(event)
	}
}

func loadRepoConfig() repoConfig {
	// Ensure the dotfiles directory exists
	repoDir := fs.NewPath(repoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating dotfiles directory", "error", err)
	}

	config := repoConfig{}
	content, err := os.ReadFile(repoConfigFilePath())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Error("error while reading dotfiles config", "error", err)
		}

		return config
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Error("invalid dotfiles config", "error", err)
	}

	return config
}
