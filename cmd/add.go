package cmd

import (
	"github.com/mlang97/dotx/app"
	"github.com/spf13/cobra"
)

func newCmdAdd(dotx app.App) *cobra.Command {
	var optDir string

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add file to your dotfiles repo",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			dotx.AddDotfile(args[0], optDir)
		},
	}

	addCmd.Flags().StringVarP(&optDir, "dir", "d", "", "test")

	return addCmd
}
