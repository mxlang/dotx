package config

import (
	"fmt"
	"path/filepath"

	"github.com/mxlang/dotx/internal/cmd"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type Event string

const (
	OnInit   Event = "init"
	OnPull   Event = "pull"
	OnDeploy Event = "deploy"
)

func (e *Event) UnmarshalYAML(unmarshal func(any) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	switch Event(value) {
	case OnInit, OnPull, OnDeploy:
		*e = Event(value)
		return nil
	default:
		return fmt.Errorf("invalid on value: %s. Must be one of: %s, %s, %s", value, OnInit, OnPull, OnDeploy)
	}
}

type runCondition string

const (
	runAlways  runCondition = "always"
	runOnce    runCondition = "once"
	runChanged runCondition = "changed"
)

func (r *runCondition) UnmarshalYAML(unmarshal func(any) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	switch runCondition(value) {
	case runAlways, runOnce, runChanged:
		*r = runCondition(value)
		return nil
	default:
		return fmt.Errorf("invalid run value: %s. Must be one of: %s, %s, %s", value, runAlways, runOnce, runChanged)
	}
}

type script struct {
	Path         string       `yaml:"path"`
	Event        Event        `yaml:"on"`
	RunCondition runCondition `yaml:"run,omitempty"`
}

type scripts []script

func (s *scripts) filter(event Event) {
	// TODO filter scripts by event
}

func (s *script) execute(event Event) {
	if s.Event != event {
		return
	}

	path := fs.NewPath(filepath.Join(repoDirPath(), s.Path))
	if !path.Exists() {
		logger.Warn("not found", "script", path.AbsPath())
		return
	}

	switch s.RunCondition {
	case runOnce:
		fmt.Println("run condition is once")
		// TODO check file already executed
	case runChanged:
		fmt.Println("run condition is changed")
		// TODO check file changed
	}

	logger.Info("execute", "script", path.AbsPath())
	if err := cmd.Run(path.AbsPath()); err != nil {
		logger.Warn("failed to execute", "script", path.AbsPath(), "error", err)
	} else {
		logger.Debug("successfully executed", "script", path.AbsPath())
	}
}
