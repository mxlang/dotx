package config

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"os"
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
	destinationPath := strings.Replace(dest.AbsPath(), repoDirPath(), "", 1)

	dotfile := Dotfile{
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
