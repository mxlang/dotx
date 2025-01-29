package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/viper"
)

type RepoConfig struct {
	Dotfiles []Dotfile
}

type Dotfile struct {
	Source      string
	Destination string
}

func FromRepoFile() RepoConfig {
	config := RepoConfig{}

	viper.SetConfigName("dotx")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.dotfiles")

	viper.ReadInConfig()
	viper.Unmarshal(&config)

	return config
}

func (r RepoConfig) WriteDotfile(dotfile Dotfile) error {
	ensureRepoConfigFile()
	if slices.Contains(r.Dotfiles, dotfile) {
		return errors.New("dotfile already added")
	}

	viper.Set("dotfiles", append(r.Dotfiles, dotfile))

	if err := viper.WriteConfig(); err != nil {
		return errors.New("failed to write repo config")
	}

	return nil
}

func ensureRepoConfigFile() {
	home, _ := os.UserHomeDir()
	repoConfigFile := filepath.Join(home, ".dotfiles", "dotx.yaml")

	if _, err := os.Stat(repoConfigFile); err != nil {
		if _, err := os.Create(repoConfigFile); err != nil {
			log.Fatal(err)
		}
	}
}
