package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
	"github.com/spf13/cobra"
	"path/filepath"
)

func newCmdDeploy(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy your dotfiles to the current system",
		Long:    "Create symbolic links from your dotfiles to their appropriate locations in your home directory",
		Example: `  dotx deploy`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			for _, dotfile := range cfg.Repo.Dotfiles {
				source := fs.NewPath(filepath.Join(cfg.RepoPath, dotfile.Source))
				dest := fs.NewPath(dotfile.Destination)

				logger.Debug("deploy", "from", source.AbsPath(), "to", dest.AbsPath())

				if dest.Exists() {
					if dest.IsSymlink() && dest.SymlinkPath() == source.AbsPath() {
						logger.Debug("dotfile already deployed with dotx", "dotfile", source.Filename())
						continue
					}

					title := "File already exists. Overwrite?"
					if dest.IsDir() {
						title = "Directory already exists. Overwrite?"
					}

					overwrite, err := tui.Confirm(
						title,
						dest.AbsPath(),
					)

					if err != nil {
						logger.Error("failed to render TUI", "error", err)
					}

					if !overwrite {
						logger.Debug("overwrite cancelled")
						continue
					}

					logger.Debug("deleting", "file", dest.AbsPath())
					if err := fs.Delete(dest); err != nil {
						logger.Error("failed to delete", "error", err)
					} else {
						logger.Debug("deleted", "path", dest.AbsPath())
					}
				}

				// Ensure parent directory exists
				dir := fs.NewPath(dest.Dir())
				if err := fs.Mkdir(dir); err != nil {
					logger.Error("could not create parent directory", "error", err)
				} else {
					logger.Debug("parent directory created or already exists", "dir", dir.AbsPath())
				}

				if err := fs.Symlink(source, dest); err != nil {
					logger.Error("failed to create symlink", "error", err)
				} else {
					logger.Debug("created symlink", "from", source.AbsPath(), "to", dest.AbsPath())
				}

				logger.Info("successfully deployed", "dotfile", source.Filename())
			}
		},
	}
}
