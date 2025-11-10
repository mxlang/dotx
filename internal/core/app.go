package core

import (
	"github.com/mxlang/dotx/internal/config"
)

type App struct {
	Config config.AppConfig
	Repo   config.RepoConfig
}

func NewApp(config config.AppConfig, repo config.RepoConfig) App {
	return App{
		Config: config,
		Repo:   repo,
	}
}
