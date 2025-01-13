package cmd

import (
	"log"
	"os"

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
	if !file.DirExists(dir) {
		log.Fatal("dotfiles directory does not exist")
	}

}

func init() {
	rootCmd.AddCommand(addCmd)
}
