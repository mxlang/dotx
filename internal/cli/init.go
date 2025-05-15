package cli

import (
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
	"github.com/spf13/cobra"
	"os"
)

func newCmdInit(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Initialize by cloning a remote dotfiles repository",
		Long:    "Set up your dotfiles environment by cloning an existing Git repository containing your configuration files",
		Example: `  dotx sync init https://github.com/username/dotfiles.git`,

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			files, err := os.ReadDir(cfg.RepoPath)
			if err != nil {
				logger.Error("failed to read content from directory", "error", err)
			}

			if len(files) > 0 {
				overwrite, err := tui.Confirm(
					"Already a git repository. Overwrite?",
					"",
				)

				if err != nil {
					logger.Error("failed to render TUI", "error", err)
				}

				if !overwrite {
					return
				}

				if err := fs.Delete(fs.NewPath(cfg.RepoPath)); err != nil {
					logger.Error("failed to delete", "error", err)
				}
			}

			url := args[0]
			if err := git.Clone(cfg.RepoPath, url); err != nil {
				logger.Error("failed to clone remote repository", "error", err)
			}

			logger.Info("successfully cloned remote repository")
		},
	}
}
