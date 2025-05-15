package main

import (
	"github.com/mxlang/dotx/internal/cli"
	"github.com/mxlang/dotx/internal/config"
)

var version = "dev"

func main() {
	cfg := config.Load()
	cli.Execute(cfg, version)
}
