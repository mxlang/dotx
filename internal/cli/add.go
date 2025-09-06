package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/core"
	"github.com/mxlang/dotx/internal/fs"
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
			for _, path := range args {
				dest := fs.NewPath(path)
				filename := dest.Filename()
				source := app.Repo.Path.Join(filename)

				dotfile := config.Dotfile{
					Source:      source,
					Destination: dest,
				}

				app.Add(dotfile, optionalDir)
			}
		},
	}

	addCmd.Flags().StringVarP(&optionalDir, "dir", "d", "", "optional directory to add dotfile to")

	return addCmd
}
