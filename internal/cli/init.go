package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
	"github.com/spf13/cobra"
)

func newCmdInit(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Initialize by cloning a remote dotfiles repository",
		Long:    "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files",
		Example: `  dotx sync init https://github.com/username/dotfiles.git`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			dir := fs.NewPath(cfg.RepoPath)
			if dir.HasSubfiles() {
				overwrite, err := tui.Confirm(
					"Directory is already a Git repository. Overwrite?",
					"",
				)

				if err != nil {
					logger.Error("failed to render TUI", "error", err)
				}

				if !overwrite {
					logger.Debug("overwrite cancelled")
					return
				}

				if err := fs.Delete(dir); err != nil {
					logger.Error("failed to delete", "error", err)
				} else {
					logger.Debug("deleted", "path", dir)
				}
			}

			url := args[0]
			if err := git.Clone(dir.AbsPath(), url); err != nil {
				logger.Error("failed to clone remote dotfiles", "error", err)
			}

			logger.Info("successfully cloned remote dotfiles")
		},
	}
}
