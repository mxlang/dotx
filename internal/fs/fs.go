package fs

import (
	"fmt"
	"os"
)

func Move(source Path, dest Path) error {
	if !source.Exists() {
		return fmt.Errorf("source path does not exist: %s", source.absPath)
	}

	if dest.Exists() {
		return fmt.Errorf("destination path already exists: %s", dest.absPath)
	}

	if err := os.Rename(source.absPath, dest.absPath); err != nil {
		return err
	}

	return nil
}

func Symlink(source Path, dest Path) error {
	if !source.Exists() {
		return fmt.Errorf("source path does not exist: %s", source.absPath)
	}

	if dest.Exists() {
		return fmt.Errorf("destination path already exists: %s", dest.absPath)
	}

	if err := os.Symlink(source.absPath, dest.absPath); err != nil {
		return err
	}

	return nil
}

func Mkdir(path Path) error {
	if err := os.MkdirAll(path.absPath, 0777); err != nil {
		return err
	}

	return nil
}

func Delete(path Path) error {
	if !path.Exists() {
		return fmt.Errorf("path does not exist: %s", path.absPath)
	}

	if err := os.Remove(path.absPath); err != nil {
		return err
	}

	return nil
}
