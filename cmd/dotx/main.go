package main

import (
	"github.com/mxlang/dotx/internal/cli"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/core"
)

var version = "dev"

func main() {
	conf := config.LoadAppConfig()
	repo := config.LoadRepoConfig()

	app := core.NewApp(conf, repo)

	cli.Execute(app, version)
}
