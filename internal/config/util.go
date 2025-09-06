package config

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/mxlang/dotx/internal/fs"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
)

func appDirPath() string { // TODO change to fs.Path
	return filepath.Join(xdg.ConfigHome, baseDir)
}

func appConfigFilePath() string { // TODO change to fs.Path
	return filepath.Join(appDirPath(), appConfigFile)
}

func repoDirPath() fs.Path {
	return fs.NewPath(xdg.DataHome, baseDir, repoDir)
}

func repoConfigFilePath() fs.Path {
	return repoDirPath().Join(repoConfigFile)
}
