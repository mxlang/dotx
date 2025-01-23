package cmd

import (
	"log"
	"log/slog"
	"os"
	"os/exec"

	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/file"
	"github.com/spf13/cobra"
)

func newCmdInit(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Clone remote dotfiles repository",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			dir := os.ExpandEnv(config.DotfilesDir)
			if file.IsDir(dir) {
				log.Fatal("dotfiles directory already exist")
			}

			command := exec.Command("git", "clone", args[0], dir)
			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			slog.Info("dotfiles repository successfully cloned from remote")
		},
	}
}
