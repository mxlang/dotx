package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mlang97/dotx/config"
	"github.com/mlang97/dotx/dotfile"
	"github.com/spf13/cobra"
)

func newCmdAdd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add file to your dotfiles repository",
		Args:  cobra.ExactArgs(1),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cfg.CheckDotfilesDir(); err != nil {
				return err
			}
			cfg.EnsureRepoConfig()
			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			absSourcePath, err := filepath.Abs(args[0])
			if err != nil {
				log.Fatal("failed to get absolut path")
			}

			fileInfo, err := os.Stat(absSourcePath)
			if err != nil {
				log.Fatal("path not found")
			}

			destPath := filepath.Join(cfg.DotfilesDir, fileInfo.Name())
			if _, err := os.Stat(destPath); err == nil {
				log.Fatal("already exist in dotfiles repository")
			}

			dotfile := dotfile.Dotfile{
				Source:      absSourcePath,
				Destination: destPath,
			}

			fmt.Println(dotfile)

			if err := dotfile.Move(); err != nil {
				log.Fatal(err)
			}

			if err := dotfile.Symlink(); err != nil {
				log.Fatal(err)
			}

			if err := cfg.AddDotfile(dotfile); err != nil {
				log.Fatal(err)
			}
		},
	}
}
