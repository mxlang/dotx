package config

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
	dataConfigFile = "data.yaml"
)

func appDirPath() string {
	return filepath.Join(xdg.ConfigHome, baseDir)
}

func appConfigFilePath() string {
	return filepath.Join(appDirPath(), appConfigFile)
}

func repoDirPath() string {
	return filepath.Join(xdg.DataHome, baseDir, repoDir)
}

func repoConfigFilePath() string {
	return filepath.Join(repoDirPath(), repoConfigFile)
}

func dataConfigFilePath() string {
	return filepath.Join(xdg.DataHome, baseDir, dataConfigFile)
}
