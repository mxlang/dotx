package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/viper"
)

var repoConfigLoader *viper.Viper

type RepoConfig struct {
	Dotfiles []Dotfile
}

type Dotfile struct {
	Source      string
	Destination string
}

func FromRepoFile(appConfig AppConfig) RepoConfig {
	config := RepoConfig{}

	repoConfigLoader.SetConfigName("dotx")
	repoConfigLoader.SetConfigType("yaml")
	repoConfigLoader.AddConfigPath(appConfig.GetRepoDir())

	repoConfigLoader.ReadInConfig()
	repoConfigLoader.Unmarshal(&config)

	return config
}

func (r RepoConfig) WriteDotfile(dotfile Dotfile) error {
	home, _ := os.UserHomeDir()
	repoConfigFile := filepath.Join(home, ".dotfiles", "dotx.yaml")

	if _, err := os.Stat(repoConfigFile); err != nil {
		if _, err := os.Create(repoConfigFile); err != nil {
			return err
		}
	}

	if slices.Contains(r.Dotfiles, dotfile) {
		return errors.New("dotfile already added")
	}

	repoConfigLoader.Set("dotfiles", append(r.Dotfiles, dotfile))

	if err := repoConfigLoader.WriteConfig(); err != nil {
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

func init() {
	repoConfigLoader = viper.New()
}
