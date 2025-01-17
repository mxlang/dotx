package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mlang97/dotx/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),

	Run: add,
}

func add(cmd *cobra.Command, args []string) {
	dir := os.ExpandEnv(viper.GetString("dir"))
	if !file.IsDir(dir) {
		log.Fatal("dotfiles directory does not exist")
	}

	dotfile := args[0]
	fileInfo, err := os.Stat(dotfile)
	if err != nil {
		log.Fatal("file does not exist")
	}

	newFilePath := filepath.Join(dir, fileInfo.Name())

	if err := os.Rename(dotfile, newFilePath); err != nil {
		log.Fatal("failed to move file")
	}

	if err := os.Symlink(newFilePath, dotfile); err != nil {
		log.Fatal("failed to symlink")
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
