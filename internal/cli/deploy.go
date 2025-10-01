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
	var force bool

	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your dotfiles to the current system",
		Long:  "Create symbolic links from your dotfiles to their appropriate locations in your home directory",
		Example: `  dotx deploy
  dotx deploy --force`,

		Args: cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
			runDeploy(cfg, force)
		},
	}

	deployCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "never prompt for overwriting")

	return deployCmd
}

func runDeploy(cfg *config.Config, force bool) {
	for _, dotfile := range cfg.Repo.Dotfiles {
		source := fs.NewPath(filepath.Join(cfg.RepoPath, dotfile.Source))
		dest := fs.NewPath(dotfile.Destination)

		logger.Debug("deploy dotfile", "from", source.AbsPath(), "to", dest.AbsPath())

		if dest.Exists() {
			if dest.IsSymlink() && dest.SymlinkPath() == source.AbsPath() {
				logger.Debug("dotfile already deployed with dotx", "dotfile", source.Filename())
				continue
			}

			if !force {
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
			}

			logger.Debug("delete", "path", dest.AbsPath())
			if err := fs.Delete(dest); err != nil {
				logger.Error("failed to delete", "error", err)
			}
		}

		// Ensure parent directory exists
		dir := fs.NewPath(dest.Dir())
		logger.Debug("create parent directory if not exists", "dir", dir.AbsPath())
		if err := fs.Mkdir(dir); err != nil {
			logger.Error("could not create parent directory", "error", err)
		}

		logger.Debug("create symlink", "from", source.AbsPath(), "to", dest.AbsPath())
		if err := fs.Symlink(source, dest); err != nil {
			logger.Error("failed to create symlink", "error", err)
		}

		logger.Info("successfully deployed", "dotfile", source.Filename())
	}

	cfg.Repo.ExecuteScripts(config.OnDeploy)
}
