package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/mlang97/dotx/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	remoteRepo string

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize empty git repo or clone remote one",

		Run: initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) {
	dir := os.ExpandEnv(viper.GetString("dir"))
	if file.IsDir(dir) {
		log.Fatal("dotfiles directory does not exist")
	}

	if remoteRepo != "" {
		command := exec.Command("git", "clone", remoteRepo, dir)
		if err := command.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		command := exec.Command("git", "init", "-b", "main", dir)
		if err := command.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&remoteRepo, "remote", "r", "", "git repository url for dotfiles")
}
