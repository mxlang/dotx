package fs

import (
	"errors"
	"os"
)

type Filesystem interface {
	Move(source File, dest File) error
	Symlink(source File, dest File) error
	Mkdir(path File) error
}

type Fs struct {
}

func NewFs() Fs {
	return Fs{}
}

func (fs Fs) Move(source File, dest File) error {
	if !source.Exists() {
		return errors.New("source path does not exist")
	}

	if dest.Exists() {
		return errors.New("destination path already exists")
	}

	if err := os.Rename(source.path, dest.path); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func (fs Fs) Symlink(source File, dest File) error {
	if !source.Exists() {
		return errors.New("source path does not exist")
	}

	if dest.Exists() {
		return errors.New("destination path already exists")
	}

	if err := os.Symlink(source.path, dest.path); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}

func (fs Fs) Mkdir(path File) error {
	return os.MkdirAll(path.path, 0777)
}
