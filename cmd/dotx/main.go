package main

import (
	"github.com/mlang97/dotx/internal/cli"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/dotx"
)

func main() {
	app := dotx.New(
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	cli.Execute(app)
}
