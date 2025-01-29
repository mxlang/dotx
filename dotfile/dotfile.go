package dotfile

import (
	"errors"
	"os"
)

type Dotfile struct {
	Source      string
	Destination string
}

func (d Dotfile) Move() error {
	if err := os.Rename(d.Source, d.Destination); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func (d Dotfile) Symlink() error {
	if err := os.Symlink(d.Destination, d.Source); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}
