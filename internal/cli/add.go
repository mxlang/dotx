package cli

import (
	"github.com/mlang97/dotx/internal/dotx"
	"github.com/spf13/cobra"
)

func newCmdAdd(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add file to your dotfiles repo",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.AddDotfile(args[0]); err != nil {
				dotx.Logger.Error("failed to add file", "error", err)
			}

			dotx.Logger.Info("successfully added to dotfiles repo")
		},
	}
}
