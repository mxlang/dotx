package git

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/mxlang/dotx/internal/config"
)

func Clone(url string) error {
	_, err := git.PlainClone(config.RepoDirPath(), false, &git.CloneOptions{
		URL: url,
	})

	return err
}

func Pull() error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	if err := worktree.Pull(&git.PullOptions{}); err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return nil
		}

		return err
	}

	return nil
}

func Add(path string) error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	_, err = worktree.Add(path)
	if err != nil {
		return err
	}

	return nil
}

func Commit(message string) error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	_, err = worktree.Commit(message, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}

func Push() error {
	repo, err := git.PlainOpen(config.RepoDirPath())
	if err != nil {
		return err
	}

	return repo.Push(&git.PushOptions{})
}
