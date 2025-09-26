package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type RepoConfig struct {
	Path fs.Path

	Dotfiles []Dotfile `yaml:"dotfiles"`
	Scripts  []script  `yaml:"scripts"`
}

func (r *RepoConfig) HasDotfile(dotfile Dotfile) bool {
	for _, dot := range r.Dotfiles {
		if dot == dotfile {
			return true
		}
	}

	return false
}

func (r *RepoConfig) AddDotfile(dotfile Dotfile) error {
	r.Dotfiles = append(r.Dotfiles, dotfile)

	config, err := yaml.Marshal(r)
	if err != nil {
		return fmt.Errorf("unable to marshal dotfiles config: %w", err)
	}

	if err := os.WriteFile(repoConfigFilePath().AbsPath(), config, 0644); err != nil {
		return fmt.Errorf("unable to write dotfiles config: %w", err)
	}

	return nil
}

func (r *RepoConfig) ExecuteScripts(e event) {
	for _, script := range r.Scripts {
		script.execute(e)
	}
}

func LoadRepoConfig() RepoConfig {
	// Ensure the dotfiles directory exists
	if err := fs.Mkdir(repoDirPath()); err != nil {
		logger.Error("error while creating dotfiles directory", "error", err)
	}

	config := RepoConfig{
		Path: repoDirPath(),
	}
	path := repoConfigFilePath().AbsPath()

	content, err := os.ReadFile(path)
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
