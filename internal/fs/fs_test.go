package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMove(t *testing.T) {
	var tests = []struct {
		test   string
		source string
		dest   string
		errMsg string

		cleanup func(fs Fs, sourcePath string, destPath string)
	}{
		{"rename .bashrc to .zshrc", filepath.Join("testdata", ".bashrc"), filepath.Join("testdata", ".zshrc"), "",
			func(fs Fs, sourcePath string, destPath string) {
				source := NewPath(sourcePath)
				dest := NewPath(destPath)
				fs.Move(dest, source)
			}},
		{"dest dir does not exist", filepath.Join("testdata", ".bashrc"), filepath.Join("test", ".bashrc"), "failed to move",
			func(fs Fs, source string, dest string) {}},
		{"source path does not exist", filepath.Join("testdata", ".zshrc"), filepath.Join("testdata", ".bashrc"), "source path does not exist",
			func(fs Fs, source string, dest string) {}},
		{"dest path already exists", filepath.Join("testdata", ".bashrc"), filepath.Join("testdata", ".bashrc"), "destination path already exists",
			func(fs Fs, source string, dest string) {}},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			fs := NewFs()

			source := NewPath(tt.source)
			dest := NewPath(tt.dest)

			err := fs.Move(source, dest)
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("got %s, want %s", err.Error(), tt.errMsg)
			}

			t.Cleanup(func() {
				tt.cleanup(fs, tt.source, tt.dest)
			})
		})
	}
}

func TestSymlink(t *testing.T) {
	var tests = []struct {
		test   string
		source string
		dest   string
		errMsg string

		cleanup func(fs Fs, sourcePath string, destPath string)
	}{
		{"create valid symlink", filepath.Join("testdata", ".bashrc"), filepath.Join("testdata", ".zshrc"), "",
			func(fs Fs, source string, dest string) {
				os.Remove(dest)
			}},
		{"source path does not exist", filepath.Join("testdata", ".zshrc"), filepath.Join("testdata", ".bashrc"), "source path does not exist",
			func(fs Fs, source string, dest string) {}},
		{"dest path already exists", filepath.Join("testdata", ".bashrc"), filepath.Join("testdata", ".bashrc"), "destination path already exists",
			func(fs Fs, source string, dest string) {}},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			fs := NewFs()

			source := NewPath(tt.source)
			dest := NewPath(tt.dest)

			err := fs.Symlink(source, dest)
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("got %s, want %s", err.Error(), tt.errMsg)
			}

			t.Cleanup(func() {
				tt.cleanup(fs, tt.source, tt.dest)
			})
		})
	}
}

func TestMkdir(t *testing.T) {
	fs := NewFs()
	path := NewPath(filepath.Join("testdata", "test"))

	err := fs.Mkdir(path)
	if err != nil {
		t.Errorf("got %s, want no error", err.Error())
	}

	os.Remove(path.absPath)
}
