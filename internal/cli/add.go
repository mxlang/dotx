package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newCmdAdd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "add <path>",
		Short: "Add a file or directory to your dotfiles",
		Long:  "Track a configuration file or directory in your dotfiles by creating a symlink to its original location",
		Example: `  dotx add ~/.bashrc
  dotx add ~/.config/nvim`,

		Args: cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			for _, path := range args {
				runAdd(cfg, path)
			}
		},
	}
}

func runAdd(cfg *config.Config, path string) {
	source := fs.NewPath(path)
	filename := source.Filename()
	dest := fs.NewPath(filepath.Join(cfg.RepoPath, filename))

	if cfg.Repo.DotfileExists(source) {
		logger.Error("already exists in dotfiles")
	}

	logger.Debug("move", "from", source.AbsPath(), "to", dest.AbsPath())
	if err := fs.Move(source, dest); err != nil {
		logger.Error("failed to move", "error", err)
	}

	logger.Debug("create symlink", "from", dest.AbsPath(), "to", source.AbsPath())
	if err := fs.Symlink(dest, source); err != nil {
		logger.Error("failed to create symlink", "error", err)
	}

	logger.Debug("write to dotfiles config")
	if err := cfg.Repo.WriteDotfile(source, dest); err != nil {
		logger.Error("failed to write dotfiles config", "error", err)
	}

	logger.Info("successfully added", "dotfile", source.Filename())
}
