package fs

import (
	"errors"
	"os"
	"path/filepath"
)

type Filesystem interface {
	AbsPath(string) (string, error)
	Move(sourcePath string, destinationPath string) error
	Symlink(sourcePath string, destinationPath string) error
	Mkdir(path string) error
}

type RepoFs struct {
}

func New() RepoFs {
	return RepoFs{}
}

func (fs RepoFs) AbsPath(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (fs RepoFs) Move(sourcePath string, destinationPath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return errors.New("source path not found")
	}

	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("already exist")
	}

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func (fs RepoFs) Symlink(sourcePath string, destinationPath string) error {
	if err := os.Symlink(sourcePath, destinationPath); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}

func (fs RepoFs) Mkdir(path string) error {
	return os.MkdirAll(path, 0777)
}
