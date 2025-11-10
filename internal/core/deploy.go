package core

import (
	"fmt"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
)

func (a App) Deploy(force bool) error {
	for _, dotfile := range a.Repo.Dotfiles {
		if err := deploy(dotfile, force); err != nil {
			return err
		}
	}

	a.Repo.ExecuteScripts(config.OnDeploy)
	return nil
}

func deploy(dotfile config.Dotfile, force bool) error {
	logger.Debug("deploy dotfile", "from", dotfile.Source, "to", dotfile.Destination)

	if !dotfile.Source.Exists() {
		logger.Warn("configured dotfile does not exist", "dotfile", dotfile.TruncateRepoPath())
		return nil
	}

	if dotfile.Destination.Exists() {
		if dotfile.Deployed() {
			logger.Debug("dotfile already deployed with dotx", "dotfile", dotfile.TruncateRepoPath())
			return nil
		}

		if !force {
			title := "File already exists. Overwrite?"
			if dotfile.Destination.IsDir() {
				title = "Directory already exists. Overwrite?"
			}

			overwrite, err := tui.Confirm(
				title,
				dotfile.Destination.AbsPath(),
			)

			if err != nil {
				return fmt.Errorf("failed to render TUI: %w", err)
			}

			if !overwrite {
				logger.Debug("overwrite cancelled")
				return nil
			}
		}

		logger.Debug("delete", "path", dotfile.Destination)
		if err := fs.Delete(dotfile.Destination); err != nil {
			return fmt.Errorf("failed to delete: %w", err)
		}
	}

	// Ensure parent directory exists
	dir := fs.NewPath(dotfile.Destination.Dir())
	logger.Debug("create parent directory if not exists", "dir", dir)
	if err := fs.Mkdir(dir); err != nil {
		return fmt.Errorf("could not create parent directory: %w", err)
	}

	logger.Debug("create symlink", "from", dotfile.Source, "to", dotfile.Destination)
	if err := fs.Symlink(dotfile.Source, dotfile.Destination); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	logger.Info("successfully deployed", "dotfile", dotfile.TruncateRepoPath())
	return nil
}
