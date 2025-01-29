package main

import (
	"github.com/mlang97/dotx/app"
	"github.com/mlang97/dotx/cmd"
	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/fs"
	"github.com/mlang97/dotx/logger"
)

func main() {
	logger := logger.New()
	fs := fs.New()
	appConfig := config.FromAppFile()
	rempoConfig := config.FromRepoFile(appConfig)

	dotx := app.New(logger, fs, appConfig, rempoConfig)

	cmd.Execute(dotx)
}
