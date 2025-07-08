package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
)

type dotfile struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

type scripts struct {
	Init []string `yaml:"init"`
}

type repoConfig struct {
	Dotfiles []dotfile `yaml:"dotfiles"`
	Scripts  scripts   `yaml:"scripts,omitempty"`
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
	// normalize paths
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
