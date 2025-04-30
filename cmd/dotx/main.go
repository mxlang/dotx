package main

import (
	"github.com/mxlang/dotx/internal/cli"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/dotx"
)

var version = "dev"

func main() {
	app := dotx.New(
		version,
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	cli.Execute(app)
}
