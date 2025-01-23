package main

import (
	"github.com/mlang97/dotx/cmd"
	"github.com/mlang97/dotx/config"
)

func main() {
	config.LoadConfig()

	cmd.Execute()
}
