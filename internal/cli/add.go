package cli

import (
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdAdd(app core.App) *cobra.Command {
	var optionalDir string

	addCmd := &cobra.Command{
		Use:   "add <path>",
		Short: "Add one or more files or directory to your dotfiles",
		Long:  "Track a configuration file or directory in your dotfiles by creating a symlink to its original location",
		Example: `  dotx add ~/.bashrc
  dotx add ~/.config/nvim
  dotx add ~/.bashrc ~/.zshrc
  dotx add -d starship ~/.config/starship.toml`,

		Args: cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			if err := app.Add(args, optionalDir); err != nil {
				logger.Error("failed to add new dotfile", "error", err)
			}
		},
	}

	addCmd.Flags().StringVarP(&optionalDir, "dir", "d", "", "optional directory to add dotfile to")

	return addCmd
}
