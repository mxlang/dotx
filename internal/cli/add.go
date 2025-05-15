package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newCmdAdd(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a file or directory to your dotfiles",
		Long:  "Track a configuration file or directory in your dotfiles by creating a symlink to its original location",
		Example: `  dotx add ~/.bashrc
  dotx add ~/.config/nvim`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			source := fs.NewPath(path)
			filename := source.Filename()
			dest := fs.NewPath(filepath.Join(cfg.RepoPath, filename))

			// check dotfile already added to repo
			if cfg.Repo.DotfileExists(source) {
				logger.Error("already exists in dotfiles")
			}

			if err := fs.Move(source, dest); err != nil {
				logger.Error("failed to move", "error", err)
			} else {
				logger.Debug("moved", "from", source.AbsPath(), "to", dest.AbsPath())
			}

			if err := fs.Symlink(dest, source); err != nil {
				logger.Error("failed to create symlink", "error", err)
			} else {
				logger.Debug("created symlink", "from", dest.AbsPath(), "to", source.AbsPath())
			}

			if err := cfg.Repo.WriteDotfile(source, dest); err != nil {
				logger.Error("failed to write dotfiles config", "error", err)
			} else {
				logger.Debug("written to dotfiles config")
			}

			logger.Info("successfully added to dotfiles", "dotfile", source.Filename())
		},
	}
}
