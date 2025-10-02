package core

import (
	"fmt"

	"github.com/mxlang/dotx/internal/config"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

func (a App) Add(paths []string, optionalDir string) error {
	for _, path := range paths {
		dest := fs.NewPath(path)
		filename := dest.Filename()
		source := a.Repo.Path.Join(filename)

		if optionalDir != "" {
			dir := a.Repo.Path.Join(optionalDir)
			if err := fs.Mkdir(dir); err != nil {
				return fmt.Errorf("could not create directory: %w", err)
			}
			source = a.Repo.Path.Join(optionalDir, filename)
		}

		dotfile := config.Dotfile{
			Source:      source,
			Destination: dest,
		}

		if err := add(a.Repo, dotfile); err != nil {
			return err
		}
	}

	return nil
}

func add(repo config.RepoConfig, dotfile config.Dotfile) error {
	if repo.HasDotfile(dotfile) {
		return fmt.Errorf("already exists in dotfiles")
	}

	logger.Debug("move", "from", dotfile.Destination, "to", dotfile.Source)
	if err := fs.Move(dotfile.Destination, dotfile.Source); err != nil {
		return fmt.Errorf("failed to move: %w", err)
	}

	logger.Debug("create symlink", "from", dotfile.Source, "to", dotfile.Destination)
	if err := fs.Symlink(dotfile.Source, dotfile.Destination); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	logger.Debug("write to dotfiles config")
	if err := repo.AddDotfile(dotfile); err != nil {
		return fmt.Errorf("failed to write dotfiles config: %w", err)
	}

	logger.Info("successfully added", "dotfile", dotfile.Source.Filename())
	return nil
}
