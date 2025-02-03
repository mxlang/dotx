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
	SymlinkPath(path string) (string, error)
}

type Fs struct {
}

func New() Fs {
	return Fs{}
}

func (fs Fs) AbsPath(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (fs Fs) Move(sourcePath string, destinationPath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		if os.IsNotExist(err) {
			return errors.New("source path does not exist")
		}
		return err
	}

	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("already exist")
	}

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func (fs Fs) Symlink(sourcePath string, destinationPath string) error {
	if err := os.Symlink(sourcePath, destinationPath); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}

func (fs Fs) Mkdir(path string) error {
	return os.MkdirAll(path, 0777)
}

func (fs Fs) SymlinkPath(path string) (string, error) {
	symlinkInfo, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("path does not exist")
		}
		return "", err
	}

	if symlinkInfo.Mode()&os.ModeSymlink != 0 {
		target, err := filepath.EvalSymlinks(path)
		if err != nil {
			return "", err
		}
		return target, nil
	}

	return "", errors.New("path is not a symlink")
}
