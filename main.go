package main

import (
	"github.com/mlang97/dotx/app"
	"github.com/mlang97/dotx/cmd"
	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/fs"
	"github.com/mlang97/dotx/log"
)

func main() {
	logger := log.New()
	fs := fs.NewFs()
	appConfig := config.FromAppFile()
	repoConfig := config.FromRepoFile(appConfig)

	dotx := app.New(logger, fs, appConfig, repoConfig)

	cmd.Execute(dotx)
}
