package dotx

import (
	"github.com/adrg/xdg"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/fs"
	"os"
	"testing"
)

const (
	xdgBaseDir          = "/tmp"
	appRootDir          = "/tmp/dotx"
	localBashrc         = "testdata/.bashrc"
	expectedSymlinkPath = "/tmp/dotx/dotfiles/.bashrc"
)

func setupTest() {
	if err := os.Setenv("XDG_CONFIG_HOME", xdgBaseDir); err != nil {
		panic(err)
	}
	if err := os.Setenv("XDG_DATA_HOME", xdgBaseDir); err != nil {
		panic(err)
	}
	xdg.Reload()
}

func cleanupTest() {
	_ = os.Remove(localBashrc)
	_, _ = os.Create(localBashrc)
	_ = os.RemoveAll(appRootDir)
}

func TestAddDotfile(t *testing.T) {
	setupTest()
	t.Cleanup(cleanupTest)

	dotx := New(
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	err := dotx.AddDotfile(localBashrc)
	if err != nil {
		t.Errorf("got error %s, want no error", err.Error())
	}

	dotfile := fs.NewPath(localBashrc)
	if !dotfile.IsSymlink() {
		t.Errorf("expected to be a symlink")
	}

	if dotfile.SymlinkPath() != expectedSymlinkPath {
		t.Errorf("got %s, want %s", dotfile.SymlinkPath(), expectedSymlinkPath)
	}
}
