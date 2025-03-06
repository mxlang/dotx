package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilename(t *testing.T) {
	var tests = []struct {
		test string
		path string
		name string
	}{
		{"filename .bashrc", filepath.Join("testdata", ".bashrc"), ".bashrc"},
		{"filename .zshrc but file not exists", filepath.Join("testdata", ".zshrc"), ".zshrc"},
		{"dirname testdata", "testdata", "testdata"},
		{"dirname test but dir not exists", "test", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			filename := path.Filename()
			if filename != tt.name {
				t.Errorf("got %s, want %s", filename, tt.name)
			}
		})
	}
}

func TestAbsPath(t *testing.T) {
	wd, _ := os.Getwd()

	var tests = []struct {
		test    string
		path    string
		absPath string
	}{
		{"file abs path exists", filepath.Join("testdata", ".bashrc"), filepath.Join(wd, "testdata", ".bashrc")},
		{"file abs path not exists", filepath.Join("testdata", ".zshrc"), filepath.Join(wd, "testdata", ".zshrc")},
		{"dir abs path exists", "testdata", filepath.Join(wd, "testdata")},
		{"dir abs path not exists", "test", filepath.Join(wd, "test")},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			absPath := path.AbsPath()
			if absPath != tt.absPath {
				t.Errorf("got %s, want %s", absPath, tt.absPath)
			}
		})
	}
}

func TestDir(t *testing.T) {
	wd, _ := os.Getwd()

	var tests = []struct {
		test string
		path string
		dir  string
	}{
		{"file path exists", filepath.Join("testdata", ".bashrc"), filepath.Join(wd, "testdata")},
		{"file path not exists", filepath.Join("testdata", ".zshrc"), filepath.Join(wd, "testdata")},
		{"dir path exists", filepath.Join("testdata"), filepath.Join(wd, "testdata")},
		{"dir path not exists", filepath.Join("test"), wd},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			dir := path.Dir()
			if dir != tt.dir {
				t.Errorf("got %s, want %s", dir, tt.dir)
			}
		})
	}
}

func TestExists(t *testing.T) {
	wd, _ := os.Getwd()

	var tests = []struct {
		test        string
		path        string
		shouldExist bool
	}{
		{"file relative exists", filepath.Join("testdata", ".bashrc"), true},
		{"file relative not exists", filepath.Join("testdata", ".zshrc"), false},
		{"file absoulte exists", filepath.Join(wd, "testdata", ".bashrc"), true},
		{"file absoulte not exists", filepath.Join(wd, "testdata", ".zshrc"), false},
		{"dir relative exists", "testdata", true},
		{"dir relative not exists", "test", false},
		{"dir absolute exists", filepath.Join(wd, "testdata"), true},
		{"dir absolute not exists", filepath.Join(wd, "test"), false},
		{"file with $HOME", filepath.Join("$HOME", "test"), false},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			exists := path.Exists()
			if exists != tt.shouldExist {
				t.Errorf("got %t, want %t", exists, tt.shouldExist)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	var tests = []struct {
		test  string
		path  string
		isDir bool
	}{
		{"dir exists", "testdata", true},
		{"dir not exists", "test", false},
		{"check existing file", filepath.Join("testdata", ".bashrc"), false},
		{"check not existing file", filepath.Join("testdata", ".zshrc"), false},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			isDir := path.IsDir()
			if isDir != tt.isDir {
				t.Errorf("got %t, want %t", isDir, tt.isDir)
			}
		})
	}
}

func TestIsSymlink(t *testing.T) {
	var tests = []struct {
		test      string
		path      string
		isSymlink bool
	}{
		{"existing symlink", filepath.Join("testdata", "symlink"), true},
		{"no symlink", filepath.Join("testdata", ".bashrc"), false},
		{"removed symlink source", filepath.Join("testdata", "symlink_removed"), true},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			isSymlink := path.IsSymlink()
			if isSymlink != tt.isSymlink {
				t.Errorf("got %t, want %t", isSymlink, tt.isSymlink)
			}
		})
	}
}

func TestSymlinkPath(t *testing.T) {
	wd, _ := os.Getwd()

	var tests = []struct {
		test        string
		path        string
		symlinkPath string
	}{
		{"existing symlink", filepath.Join("testdata", "symlink"), filepath.Join(wd, "testdata", ".bashrc")},
		{"no symlink", filepath.Join("testdata", ".bashrc"), ""},
		{"removed symlink source", filepath.Join("testdata", "symlink_removed"), ""},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			path := NewPath(tt.path)
			symlinkPath := path.SymlinkPath()
			if symlinkPath != tt.symlinkPath {
				t.Errorf("got %s, want %s", symlinkPath, tt.symlinkPath)
			}
		})
	}
}
