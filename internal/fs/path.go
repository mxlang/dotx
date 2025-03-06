package fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	absPath string
}

func NewPath(path string) Path {
	return Path{
		absPath: normalizePath(path),
	}
}

func (p Path) Filename() string {
	if p.Exists() {
		fileInfo, _ := os.Stat(p.absPath)
		return fileInfo.Name()
	}

	return filepath.Base(p.absPath)
}

func (p Path) AbsPath() string {
	return p.absPath
}

func (p Path) Dir() string {
	if p.IsDir() {
		return p.absPath
	}

	return filepath.Dir(p.absPath)
}

func (p Path) Exists() bool {
	_, err := os.Stat(p.absPath)
	return err == nil || !errors.Is(err, fs.ErrNotExist)
}

func (p Path) IsDir() bool {
	if p.Exists() {
		fileInfo, _ := os.Stat(p.absPath)
		return fileInfo.IsDir()
	}

	return false
}

func (p Path) IsSymlink() bool {
	symlinkInfo, err := os.Lstat(p.absPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	if symlinkInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}

	return false
}

func (p Path) SymlinkPath() string {
	if p.IsSymlink() {
		symlink, _ := filepath.EvalSymlinks(p.absPath)
		return symlink
	}

	return ""
}

func normalizePath(path string) string {
	if strings.Contains(path, "$HOME") {
		path = os.ExpandEnv(path)
	}

	if filepath.IsAbs(path) {
		path = filepath.Clean(path)
	}

	path, _ = filepath.Abs(path)

	return path
}
