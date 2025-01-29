package config

import (
	"errors"
	"os"
	"path/filepath"
)

type Dotfile struct {
	Destination string
	Source      string
}

func New(appConfig appConfig, path string) (Dotfile, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return Dotfile{}, errors.New("failed to get absolut path")
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return Dotfile{}, errors.New("path not found")
	}

	destPath := filepath.Join(appConfig.DotfilesDir, fileInfo.Name())
	if _, err := os.Stat(destPath); err == nil {
		return Dotfile{}, errors.New("already exist in dotfiles repository")
	}

	dotfile := Dotfile{
		Destination: destPath,
		Source:      absPath,
	}

	return dotfile, nil
}

func (d Dotfile) move() error {
	if err := os.Rename(d.Source, d.Destination); err != nil {
		return errors.New("failed to move")
	}

	return nil
}

func (d Dotfile) symlink() error {
	if err := os.Symlink(d.Destination, d.Source); err != nil {
		return errors.New("failed to symlink")
	}

	return nil
}
