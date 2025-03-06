package fs

import (
	"errors"
	"os"
)

type Filesystem interface {
	Move(source Path, dest Path) error
	Symlink(source Path, dest Path) error
	Mkdir(path Path) error
}

type Fs struct {
}

func NewFs() Fs {
	return Fs{}
}

func (fs Fs) Move(source Path, dest Path) error {
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

func (fs Fs) Symlink(source Path, dest Path) error {
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

func (fs Fs) Mkdir(path Path) error {
	return os.MkdirAll(path.absPath, 0777)
}
