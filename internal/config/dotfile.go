package config

import (
	"os"
	"strings"

	"github.com/mxlang/dotx/internal/fs"
)

type Dotfile struct {
	Source      fs.Path `yaml:"source"`
	Destination fs.Path `yaml:"destination"`
}

func (d *Dotfile) UnmarshalYAML(unmarshal func(any) error) error {
	var temp struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	}

	if err := unmarshal(&temp); err != nil {
		return err
	}

	d.Source = repoDirPath().Join(temp.Source)
	d.Destination = fs.NewPath(temp.Destination)

	return nil
}

func (d Dotfile) MarshalYAML() (any, error) {
	type temp struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	}

	home, _ := os.UserHomeDir()
	a := temp{
		Source:      d.TruncateRepoPath(),
		Destination: strings.Replace(d.Destination.AbsPath(), home, "$HOME", 1),
	}

	return a, nil
}

func (d Dotfile) Deployed() bool {
	return d.Destination.IsSymlink() && d.Destination.SymlinkPath() == d.Source.AbsPath()
}

func (d Dotfile) TruncateRepoPath() string {
	return strings.Replace(d.Source.AbsPath(), repoDirPath().AbsPath(), "", 1)
}
