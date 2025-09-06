package git

import (
	"errors"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/mxlang/dotx/internal/fs"

	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func Clone(repoDir fs.Path, url string) error {
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

	_, err := git.PlainClone(repoDir.AbsPath(), gitOpts)
	if err != nil {
		return err
	}

	return nil
}

func Pull(repoDir fs.Path) error {
	repo, err := git.PlainOpen(repoDir.AbsPath())
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

func Add(repoDir fs.Path, path string) error {
	repo, err := git.PlainOpen(repoDir.AbsPath())
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

func Commit(repoDir fs.Path, message string) error {
	repo, err := git.PlainOpen(repoDir.AbsPath())
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

func Push(repoDir fs.Path) error {
	repo, err := git.PlainOpen(repoDir.AbsPath())
	if err != nil {
		return err
	}

	if err := repo.Push(&git.PushOptions{}); err != nil {
		return err
	}

	return nil
}

func Remote(repoDir fs.Path) ([]string, error) {
	repo, err := git.PlainOpen(repoDir.AbsPath())
	if err != nil {
		return nil, err
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		return nil, err
	}

	return remote.Config().URLs, nil
}

func IsBehindRemote(repoDir fs.Path) (bool, error) {
	repo, err := git.PlainOpen(repoDir.AbsPath())
	if err != nil {
		return false, err
	}

	err = repo.Fetch(&git.FetchOptions{RemoteName: "origin", Progress: nil, Tags: git.NoTags})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return false, err
	}

	headRef, err := repo.Head()
	if err != nil {
		return false, err
	}

	remoteRef, err := repo.Reference(plumbing.NewRemoteReferenceName("origin", headRef.Name().Short()), true)
	if err != nil {
		return false, err
	}

	headCommit, err := repo.CommitObject(headRef.Hash())
	if err != nil {
		return false, err
	}

	remoteCommit, err := repo.CommitObject(remoteRef.Hash())
	if err != nil {
		return false, err
	}

	isAncestor, err := headCommit.IsAncestor(remoteCommit)
	if err != nil {
		return false, err
	}

	return isAncestor, nil
}
