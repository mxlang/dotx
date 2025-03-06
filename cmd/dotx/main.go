package main

import (
	"github.com/mlang97/dotx/internal/cli"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/dotx"
	"github.com/mlang97/dotx/internal/fs"
)

func main() {
	fs := fs.NewFs()
	appConfig := config.FromAppFile()
	repoConfig := config.FromRepoFile(appConfig)

	dotx := dotx.New(fs, appConfig, repoConfig)

	cli.Execute(dotx)
}
