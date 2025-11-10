package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func createFile(t *testing.T, dir string, name string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func createSymlink(t *testing.T, source string, dest string) string {
	t.Helper()
	if err := os.Symlink(source, dest); err != nil {
		t.Fatal(err)
	}
	return dest
}

func TestPath_Filename(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "get filename from file",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: ".bashrc",
		},
		{
			name:     "get filename from non existing file",
			path:     filepath.Join(tempDir, ".zshrc"),
			expected: ".zshrc",
		},
		{
			name:     "get filename from directory",
			path:     tempDir,
			expected: filepath.Base(tempDir),
		},
		{
			name:     "get filename from non existing directory",
			path:     filepath.Join(tempDir, "config"),
			expected: "config",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := NewPath(tt.path)
			got := file.Filename()
			if tt.expected != got {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestPath_AbsPath(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "get absolute path from file",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: filepath.Join(tempDir, ".bashrc"),
		},
		{
			name:     "get absolute path from non existing file",
			path:     filepath.Join(tempDir, ".zshrc"),
			expected: filepath.Join(tempDir, ".zshrc"),
		},
		{
			name:     "get absolute path from directory",
			path:     tempDir,
			expected: tempDir,
		},
		{
			name:     "get absolute path from non existing directory",
			path:     filepath.Join(tempDir, "config"),
			expected: filepath.Join(tempDir, "config"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := NewPath(tt.path)
			got := file.AbsPath()
			if tt.expected != got {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestPath_Dir(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "get dir from file",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: tempDir,
		},
		{
			name:     "get dir from non existing file",
			path:     filepath.Join(tempDir, ".zshrc"),
			expected: tempDir,
		},
		{
			name:     "get dir from directory",
			path:     tempDir,
			expected: tempDir,
		},
		{
			name:     "get dir from non existing directory",
			path:     filepath.Join(tempDir, "config"),
			expected: tempDir,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := NewPath(tt.path)
			got := file.Dir()
			if tt.expected != got {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestPath_Exists(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "existing file",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: true,
		},
		{
			name:     "non existing file",
			path:     filepath.Join(tempDir, ".zshrc"),
			expected: false,
		},
		{
			name:     "existing directory",
			path:     tempDir,
			expected: true,
		},
		{
			name:     "non existing directory",
			path:     filepath.Join(tempDir, "config"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := NewPath(tt.path)
			got := path.Exists()
			if tt.expected != got {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestPath_IsDir(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "is a directory",
			path:     tempDir,
			expected: true,
		},
		{
			name:     "is not a directory",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: false,
		},
		{
			name:     "non existing directory",
			path:     filepath.Join(tempDir, "config"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := NewPath(tt.path)
			got := path.IsDir()
			if tt.expected != got {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestPath_IsSymlink(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "is a symlink to file",
			path:     createSymlink(t, createFile(t, tempDir, ".bashrc"), filepath.Join(tempDir, ".bashrc_sym")),
			expected: true,
		},
		{
			name:     "is a symlink to directory",
			path:     createSymlink(t, tempDir, filepath.Join(tempDir, "config")),
			expected: true,
		},
		{
			name:     "is not a symlink (regular file)",
			path:     createFile(t, tempDir, ".bashrc"),
			expected: false,
		},
		{
			name:     "is not a symlink (directory)",
			path:     tempDir,
			expected: false,
		},
		{
			name:     "non existing path",
			path:     filepath.Join(tempDir, "nonexistent"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := NewPath(tt.path)
			got := path.IsSymlink()
			if tt.expected != got {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestPath_SymlinkPath(t *testing.T) {
	tempDir := t.TempDir()

	targetFile := createFile(t, tempDir, ".bashrc")
	symlinkFile := createSymlink(t, targetFile, filepath.Join(tempDir, ".bashrc_sym"))

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "get symlink target path",
			path:     symlinkFile,
			expected: targetFile,
		},
		{
			name:     "not a symlink returns empty string",
			path:     targetFile,
			expected: "",
		},
		{
			name:     "directory is not a symlink",
			path:     tempDir,
			expected: "",
		},
		{
			name:     "non existing path returns empty string",
			path:     filepath.Join(tempDir, "nonexistent"),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := NewPath(tt.path)
			got := path.SymlinkPath()
			if tt.expected != got {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestPath_Join(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		basePath string
		paths    []string
		expected string
	}{
		{
			name:     "join single path",
			basePath: tempDir,
			paths:    []string{"config"},
			expected: filepath.Join(tempDir, "config"),
		},
		{
			name:     "join multiple paths",
			basePath: tempDir,
			paths:    []string{"config", "app", "settings.json"},
			expected: filepath.Join(tempDir, "config", "app", "settings.json"),
		},
		{
			name:     "join with empty paths",
			basePath: tempDir,
			paths:    []string{},
			expected: tempDir,
		},
		{
			name:     "join with dot paths",
			basePath: tempDir,
			paths:    []string{".", "config", "..", "other"},
			expected: filepath.Join(tempDir, "other"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := NewPath(tt.basePath)
			got := path.Join(tt.paths...).AbsPath()
			if tt.expected != got {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
