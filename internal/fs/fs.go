package fs

import (
	"errors"
	"os"
)

func Move(source Path, dest Path) error {
	if !source.Exists() {
		return errors.New("source path does not exist")
	}

	if dest.Exists() {
		return errors.New("destination path already exists")
	}

	if err := os.Rename(source.absPath, dest.absPath); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func Symlink(source Path, dest Path) error {
	if !source.Exists() {
		return errors.New("source path does not exist")
	}

	if dest.Exists() {
		return errors.New("destination path already exists")
	}

	if err := os.Symlink(source.absPath, dest.absPath); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}

func Mkdir(path Path) error {
	if err := os.MkdirAll(path.absPath, 0777); err != nil {
		return errors.New("failed to create directory")
	}

	return nil
}

func Delete(path Path) error {
	return os.Remove(path.absPath)
}
