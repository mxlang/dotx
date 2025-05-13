package git

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/mxlang/dotx/internal/config"
)

func Clone(url string) error {
	_, err := git.PlainClone(config.RepoDirPath(), false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return fmt.Errorf("failed to clone remote repository: %w", err)
	}

	return nil
}

func Pull() error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	if err := worktree.Pull(&git.PullOptions{}); err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return nil
		}

		return fmt.Errorf("failed to pull repository: %w", err)
	}

	return nil
}

func Add(path string) error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = worktree.Add(path)
	if err != nil {
		return fmt.Errorf("failed to add path %s: %w", path, err)
	}

	return nil
}

func Commit(message string) error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = worktree.Commit(message, &git.CommitOptions{})
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	return nil
}

func Push() error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	if err := repo.Push(&git.PushOptions{}); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	return nil
}
