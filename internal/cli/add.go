package cli

import (
	"github.com/mxlang/dotx/internal/dotx"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdAdd(dotx dotx.App) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a file or directory to your dotfiles repository",
		Long:  "Track a configuration file or directory in your dotfiles repository by creating a symlink to its original location",
		Example: `  dotx add ~/.bashrc
  dotx add ~/.config/nvim`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := dotx.AddDotfile(args[0]); err != nil {
				logger.Error("failed to add dotfile", "error", err)
			}

			logger.Info("successfully added to dotfiles repository")
		},
	}
}
