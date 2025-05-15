package config

import (
	"github.com/adrg/xdg"
	"path/filepath"
)

const (
	baseDir        = "dotx"
	appConfigFile  = "config.yaml"
	repoDir        = "dotfiles"
	repoConfigFile = "dotx.yaml"
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
