package git

import (
	"errors"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func Clone(repoDir string, url string) error {
	gitOpts := &git.CloneOptions{
		URL: url,
	}

	if strings.HasPrefix(url, "ssh://") || strings.HasPrefix(url, "git@") {
		authMethod, err := ssh.NewSSHAgentAuth("git")
		if err != nil {
			return err
		}

		gitOpts.Auth = authMethod
	}

	_, err := git.PlainClone(repoDir, gitOpts)
	if err != nil {
		return err
	}

	return nil
}

func Pull(repoDir string) error {
	repo, err := git.PlainOpen(repoDir)
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

func Add(repoDir string, path string) error {
	repo, err := git.PlainOpen(repoDir)
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

func Commit(repoDir string, message string) error {
	repo, err := git.PlainOpen(repoDir)
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

func Push(repoDir string) error {
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		return err
	}

	if err := repo.Push(&git.PushOptions{}); err != nil {
		return err
	}

	return nil
}

func Remote(repoDir string) ([]string, error) {
	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		return nil, err
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return nil, err
	}

	return remote.Config().URLs, nil
}
