package main

import (
	"github.com/mlang97/dotx/cmd"
	"github.com/mlang97/dotx/internal/app"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/fs"
	"github.com/mlang97/dotx/internal/log"
)

func main() {
	logger := log.New()
	fs := fs.NewFs()
	appConfig := config.FromAppFile()
	repoConfig := config.FromRepoFile(appConfig)

	dotx := app.New(logger, fs, appConfig, repoConfig)

	cmd.Execute(dotx)
}
