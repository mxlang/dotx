package config

import (
	"github.com/mlang97/dotx/internal/fs"
	"github.com/mlang97/dotx/internal/logger"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
)

func EnsureDirs() {
	appDir := fs.NewPath(appDirPath())
	if err := fs.Mkdir(appDir); err != nil {
		logger.Error("error while creating dotx config dir", "error", err)
	} else {
		logger.Debug("dotx config dir created or already existent", "dir", appDir)
	}

	repoDir := fs.NewPath(repoDirPath())
	if err := fs.Mkdir(repoDir); err != nil {
		logger.Error("error while creating dotfiles repo dir", "error", err)
	} else {
		logger.Debug("dotfiles repo dir created or already existent", "dir", repoDir)
	}
}
