package config

import (
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/spf13/viper"
)

var repoConfigLoader *viper.Viper

type RepoConfig struct {
	appConfig AppConfig

	Dotfiles []Dotfile
}

type Dotfile struct {
	Source      string
	Destination string
}

func FromRepoFile(appConfig AppConfig) RepoConfig {
	config := RepoConfig{
		appConfig: appConfig,
	}

	repoConfigLoader.SetConfigName("dotx")
	repoConfigLoader.SetConfigType("yaml")
	repoConfigLoader.AddConfigPath(os.ExpandEnv(appConfig.RepoDir))

	repoConfigLoader.ReadInConfig()
	repoConfigLoader.Unmarshal(&config)

	return config
}

func (r RepoConfig) GetDotfile(destPath string) Dotfile {
	for _, df := range r.Dotfiles {
		if os.ExpandEnv(df.Destination) == destPath {
			return df
		}
	}

	return Dotfile{}
}

func (r RepoConfig) WriteDotfile(dotfile Dotfile) error {
	repoConfigFile := filepath.Join(os.ExpandEnv(r.appConfig.RepoDir), "dotx.yaml")

	// create dotx.yaml in repo dir if not exists
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

func init() {
	repoConfigLoader = viper.New()
}
