package main

import (
	"github.com/mxlang/dotx/internal/cli"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/core"
)

var version = "dev"

func main() {
	repo := config.LoadRepoConfig()

	app := core.NewApp(repo)

	cli.Execute(app, version)
}
