package fs

import (
	"errors"
	iofs "io/fs"
	"os"
	"path/filepath"
)

type Path struct {
	absPath string
}

func NewPath(path ...string) Path {
	return Path{
		absPath: expandPath(filepath.Join(path...)),
	}
}

func (p Path) Filename() string {
	if info, err := os.Stat(p.absPath); err == nil {
		return info.Name()
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
	return err == nil || !errors.Is(err, iofs.ErrNotExist)
}

func (p Path) IsDir() bool {
	info, err := os.Stat(p.absPath)
	if err != nil {
		return false
	}

	return info.IsDir()
}

func (p Path) IsSymlink() bool {
	info, err := os.Lstat(p.absPath)
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}

func (p Path) SymlinkPath() string {
	if !p.IsSymlink() {
		return ""
	}
	symlink, err := filepath.EvalSymlinks(p.absPath)
	if err != nil {
		return ""
	}

	return symlink
}

func (p Path) Join(path ...string) Path {
	combinedPath := filepath.Join(append([]string{p.absPath}, path...)...)
	return NewPath(combinedPath)
}

func (p Path) String() string {
	return p.absPath
}

func expandPath(path string) string {
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
