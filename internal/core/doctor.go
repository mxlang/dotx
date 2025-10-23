package core

import (
	"github.com/charmbracelet/huh"
	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
)

func (a App) Doctor(force bool) error {
	var missingDotfiles []config.Dotfile
	for _, dotfile := range a.Repo.Dotfiles {
		if dotfile.Deployed() {
			logger.Debug("dotfile already deployed with dotx", "dotfile", dotfile.TruncateRepoPath())
		} else {
			missingDotfiles = append(missingDotfiles, dotfile)
		}
	}

	if len(missingDotfiles) == 0 {
		return nil
	}

	selected, err := tui.MultiSelect(
		"Unlinked dotfiles detected",
		"The following dotfiles exist in your repository but are not deployed by dotx. Select which files you would like to deploy:",
		missingDotfiles,
		func(t config.Dotfile) huh.Option[config.Dotfile] {
			return huh.NewOption(t.TruncateRepoPath(), t)
		})
	if err != nil {
		return err
	}

	for _, dotfile := range selected {
		err := deploy(dotfile, force)
		if err != nil {
			logger.Warn("failed to deploy dotfile", "dotfile", dotfile.TruncateRepoPath(), "error", err)
		}
	}

	return nil
}
