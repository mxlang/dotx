package cli

import (
	"path/filepath"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/spf13/cobra"
)

func newCmdAdd(cfg *config.Config) *cobra.Command {
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
				runAdd(cfg, path, optionalDir)
			}
		},
	}

	addCmd.Flags().StringVarP(&optionalDir, "dir", "d", "", "optional directory to add dotfile to")

	return addCmd
}

func runAdd(cfg *config.Config, path string, optionalDir string) {
	source := fs.NewPath(path)
	filename := source.Filename()
	dest := fs.NewPath(filepath.Join(cfg.RepoPath, filename))

	if optionalDir != "" {
		dir := fs.NewPath(filepath.Join(cfg.RepoPath, optionalDir))
		if err := fs.Mkdir(dir); err != nil {
			logger.Error("could not create directory", "dir", dir, "error", err)
		}
		dest = fs.NewPath(filepath.Join(cfg.RepoPath, optionalDir, filename))
	}

	if cfg.Repo.HasDotfile(source) {
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
	if err := cfg.Repo.AddDotfile(source, dest); err != nil {
		logger.Error("failed to write dotfiles config", "error", err)
	}

	logger.Info("successfully added", "dotfile", source.Filename())
}
