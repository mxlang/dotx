package dotx

import (
	"github.com/adrg/xdg"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/fs"
	"os"
	"path/filepath"
	"testing"
)

const (
	baseDir     = "testdata"
	appDir      = "testdata/dotx"
	localBashrc = "testdata/.bashrc"
	repoBashrc  = "testdata/dotx/dotfiles/.bashrc"
)

func setupTest() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	xdgBaseDir := filepath.Join(wd, baseDir)
	if err := os.Setenv("XDG_CONFIG_HOME", xdgBaseDir); err != nil {
		panic(err)
	}
	if err := os.Setenv("XDG_DATA_HOME", xdgBaseDir); err != nil {
		panic(err)
	}
	xdg.Reload()
}

func cleanupTest() {
	os.RemoveAll(appDir)

	// restore local bashrc
	os.Remove(localBashrc)
	os.Create(localBashrc)
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

	expectedDotfile := fs.NewPath(repoBashrc)
	if dotfile.SymlinkPath() != expectedDotfile.AbsPath() {
		t.Errorf("got %s, want %s", dotfile.SymlinkPath(), expectedDotfile.AbsPath())
	}
}

func TestAddDotfileAlreadyExists(t *testing.T) {
	setupTest()
	t.Cleanup(cleanupTest)

	dotx := New(
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	wantedErrMsg := "dotfile already exists"

	err := dotx.AddDotfile(localBashrc)
	if err != nil {
		t.Errorf("got error %s, want none", err.Error())
	}

	err = dotx.AddDotfile(localBashrc)
	if err == nil {
		t.Errorf("want error, got none")
	} else if err.Error() != wantedErrMsg {
		t.Errorf("got error %s, want %s", err.Error(), wantedErrMsg)
	}
}

func TestAddDotfileMissingSourcePath(t *testing.T) {
	setupTest()
	t.Cleanup(cleanupTest)

	dotx := New(
		config.LoadAppConfig(),
		config.LoadRepoConfig(),
	)

	wantedErrMsg := "source path does not exist"

	err := dotx.AddDotfile("testdata/.zshrc")
	if err == nil {
		t.Errorf("want error, got none")
	} else if err.Error() != wantedErrMsg {
		t.Errorf("got error %s, want %s", err.Error(), wantedErrMsg)
	}
}
