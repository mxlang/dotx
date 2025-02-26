package fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	path string
}

func NewFile(path string) File {
	return File{
		path: absPath(path),
	}
}

func (f File) Name() string {
	if f.Exists() {
		fileInfo, _ := os.Stat(f.path)
		return fileInfo.Name()
	}

	return filepath.Base(f.path)
}

func (f File) AbsPath() string {
	return f.path
}

func (f File) Dir() string {
	if f.IsDir() {
		return f.path
	}

	return filepath.Dir(f.path)
}

func (f File) Exists() bool {
	_, err := os.Stat(f.path)
	return err == nil || !errors.Is(err, fs.ErrNotExist)
}

func (f File) IsDir() bool {
	if f.Exists() {
		fileInfo, _ := os.Stat(f.path)
		return fileInfo.IsDir()
	}

	return false
}

func (f File) IsSymlink() bool {
	symlinkInfo, err := os.Lstat(f.path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	if symlinkInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}

	return false
}

func (f File) SymlinkPath() string {
	if f.IsSymlink() {
		symlink, _ := filepath.EvalSymlinks(f.path)
		return symlink
	}

	return ""
}

func absPath(path string) string {
	if strings.Contains(path, "$HOME") {
		path = os.ExpandEnv(path)
	}

	if filepath.IsAbs(path) {
		path = filepath.Clean(path)
	}

	path, _ = filepath.Abs(path)

	return path
}
