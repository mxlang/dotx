package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/mlang97/dotx/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Clone remote dotfiles repository",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			dir := os.ExpandEnv(viper.GetString("dir"))
			if file.IsDir(dir) {
				log.Fatal("dotfiles directory already exist")
			}

			command := exec.Command("git", "clone", args[0], dir)
			if err := command.Run(); err != nil {
				log.Fatal(err)
			}
		},
	}
}
