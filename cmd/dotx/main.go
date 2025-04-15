package main

import (
	"github.com/mxlang/dotx/internal/cli"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/dotx"
)

func main() {
	app := dotx.New(
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	cli.Execute(app)
}
