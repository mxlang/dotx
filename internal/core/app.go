package core

import (
	"github.com/mxlang/dotx/internal/config"
)

type App struct {
	Repo config.RepoConfig
}

func NewApp(repo config.RepoConfig) App {
	return App{
		Repo: repo,
	}
}
