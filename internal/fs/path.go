package fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
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

func (p Path) HasSubfiles() bool {
	if !p.IsDir() {
		return false
	}

	files, err := os.ReadDir(p.absPath)
	if err != nil {
		return false
	}

	return len(files) > 0
}

func normalizePath(path string) string {
	// Expand all environment variables in the path
	path = os.ExpandEnv(path)

	// Clean the path to remove any unnecessary elements
	path = filepath.Clean(path)

	// Convert to an absolute path
	absPath, err := filepath.Abs(path)
	if err == nil {
		path = absPath
	}

	return path
}
