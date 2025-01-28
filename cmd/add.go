package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mlang97/dotx/config"
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
			absFilePath, err := filepath.Abs(args[0])
			if err != nil {
				log.Fatal("failed to get absolut path for file")
			}

			fileInfo, err := os.Stat(absFilePath)
			if err != nil {
				log.Fatal("file not found")
			}

			destFilePath := filepath.Join(cfg.DotfilesDir, fileInfo.Name())
			if _, err := os.Stat(destFilePath); err == nil {
				log.Fatal("file already exist")
			}

			if err := os.Rename(absFilePath, destFilePath); err != nil {
				log.Fatal("failed to move file")
			}

			if err := os.Symlink(destFilePath, absFilePath); err != nil {
				log.Fatal("failed to symlink")
			}

			cfg.AddDotfile(absFilePath, destFilePath)
		},
	}
}
