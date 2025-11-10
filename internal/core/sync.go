package core

import (
	"fmt"
	"slices"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/git"
	"github.com/mxlang/dotx/internal/logger"
	"github.com/mxlang/dotx/internal/tui"
)

func (a App) Init(url string, deploy bool, force bool) error {
	if shouldCloneDotfiles(a.Repo.Path, url) {
		logger.Debug("clone remote dotfiles", "url", url)
		if err := git.Clone(a.Repo.Path, url); err != nil {
			return fmt.Errorf("failed to clone remote dotfiles: %w", err)
		}

		a.Repo = config.LoadRepoConfig()

		logger.Info("successfully cloned remote dotfiles")
	}

	a.Repo.ExecuteScripts(config.OnInit)

	if deploy {
		logger.Debug("automatic deploy is active")
		if err := a.Deploy(force); err != nil {
			return fmt.Errorf("failed to deploy: %w", err)
		}
	}

	return nil
}

func (a App) Pull(deploy bool, force bool) error {
	logger.Debug("pull changes from remote dotfiles")
	if err := git.Pull(a.Repo.Path); err != nil {
		return fmt.Errorf("failed to pull remote dotfiles: %w", err)
	}

	logger.Info("successfully pulled from remote dotfiles")

	a.Repo.ExecuteScripts(config.OnPull)

	if deploy {
		logger.Debug("automatic deploy is active")
		if err := a.Deploy(force); err != nil {
			return fmt.Errorf("failed to deploy: %w", err)
		}
	}

	return nil
}

func (a App) Push(commitMessage string) error {
	logger.Debug("add changes to dotfiles")
	if err := git.Add(a.Repo.Path, "."); err != nil {
		return fmt.Errorf("failed to add changes: %w", err)
	}

	if commitMessage == "" {
		msg, err := tui.Text("Write your commit message")
		if err != nil {
			return fmt.Errorf("failed to render TUI: %w", err)
		}

		commitMessage = msg
	}

	logger.Debug("commit changes to dotfiles", "message", commitMessage)
	if err := git.Commit(a.Repo.Path, commitMessage); err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	logger.Debug("push changes to dotfiles")
	if err := git.Push(a.Repo.Path); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	logger.Info("successfully pushed changes to remote dotfiles")
	return nil
}

func shouldCloneDotfiles(dir fs.Path, url string) bool {
	remotes, err := git.Remote(dir)
	if err != nil {
		logger.Debug("no remote dotfiles found")
		return true
	}

	if !slices.Contains(remotes, url) {
		overwrite, err := tui.Confirm(
			"Directory is already another Git repository. Overwrite?",
			"",
		)

		if err != nil {
			logger.Error("failed to render TUI", "error", err)
		}

		if !overwrite {
			logger.Debug("overwrite cancelled")
			return false
		}

		logger.Debug("delete", "path", dir)
		if err := fs.Delete(dir); err != nil {
			logger.Error("failed to delete", "error", err)
		}

		return overwrite
	}

	return false
}
