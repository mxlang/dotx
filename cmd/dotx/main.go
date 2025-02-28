package main

import (
	"github.com/mlang97/dotx/internal/cmd"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/dotx"
	"github.com/mlang97/dotx/internal/fs"
	"github.com/mlang97/dotx/internal/log"
)

func main() {
	logger := log.New()
	fs := fs.NewFs()
	appConfig := config.FromAppFile()
	repoConfig := config.FromRepoFile(appConfig)

	dotx := dotx.New(logger, fs, appConfig, repoConfig)

	cmd.Execute(dotx)
}
