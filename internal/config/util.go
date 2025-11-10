package config

import (
	"github.com/adrg/xdg"
	"github.com/mxlang/dotx/internal/fs"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
	dataConfigFile = "data.yaml"
)

func appDirPath() fs.Path {
	return fs.NewPath(xdg.ConfigHome, baseDir)
}

func appConfigFilePath() fs.Path {
	return appDirPath().Join(appConfigFile)
}

func repoDirPath() fs.Path {
	return fs.NewPath(xdg.DataHome, baseDir, repoDir)
}

func repoConfigFilePath() fs.Path {
	return repoDirPath().Join(repoConfigFile)
}

func dataConfigFilePath() fs.Path {
	return fs.NewPath(xdg.DataHome, baseDir, dataConfigFile)
}
