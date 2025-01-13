package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	remoteRepo string

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",

		Run: initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) {
	dir := os.ExpandEnv(dotfilesDir)

	info, err := os.Stat(dir)
	if err == nil {
		if info.IsDir() {
			log.Fatal("dotfiles repository already initialized")
		}
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
