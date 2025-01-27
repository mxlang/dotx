package cmd

import (
	"log"
	"log/slog"
	"os/exec"

	"github.com/mlang97/dotx/config"
	"github.com/spf13/cobra"
)

func newCmdInit(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Clone remote dotfiles repository",
		Args:  cobra.ExactArgs(1),

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := cfg.CheckDotfilesDir(); err == nil {
				log.Fatal("dotfiles directory already exist")
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			command := exec.Command("git", "clone", args[0], cfg.DotfilesDir)
			if err := command.Run(); err != nil {
				log.Fatal(err)
			}

			slog.Info("dotfiles repository successfully cloned from remote")
		},
	}
}
