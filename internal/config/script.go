package config

import (
	"fmt"
	"strings"

	"github.com/mxlang/dotx/internal/cmd"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type event string

const (
	OnInit   event = "init"
	OnPull   event = "pull"
	OnDeploy event = "deploy"
)

func (e *event) UnmarshalYAML(unmarshal func(any) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	switch event(value) {
	case OnInit, OnPull, OnDeploy:
		*e = event(value)
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
	Path         fs.Path      `yaml:"path"`
	Event        event        `yaml:"on"`
	RunCondition runCondition `yaml:"run,omitempty"`
	// TODO maybe add before and after hook
}

func (s *script) UnmarshalYAML(unmarshal func(any) error) error {
	var temp struct {
		Path         string       `yaml:"path"`
		Event        event        `yaml:"on"`
		RunCondition runCondition `yaml:"run"`
	}

	if err := unmarshal(&temp); err != nil {
		return err
	}

	if temp.Event == OnInit {
		if temp.RunCondition != "" {
			return fmt.Errorf("run property is not allowed for on init scripts")
		}
		s.Path = repoDirPath().Join(temp.Path)
		s.Event = temp.Event
		return nil
	}

	if temp.RunCondition == "" {
		temp.RunCondition = runAlways
	}

	s.Path = repoDirPath().Join(temp.Path)
	s.Event = temp.Event
	s.RunCondition = temp.RunCondition

	return nil
}

func (s script) MarshalYAML() (any, error) {
	type temp struct {
		Path         string       `yaml:"path"`
		Event        event        `yaml:"on"`
		RunCondition runCondition `yaml:"run,omitempty"`
	}

	out := temp{
		Path:  strings.Replace(s.Path.AbsPath(), repoDirPath().AbsPath(), "", 1),
		Event: s.Event,
	}

	// For init scripts, never marshal `run` because it's implicitly `once` and not configurable
	if s.Event != OnInit && s.RunCondition != "" && s.RunCondition != runAlways {
		out.RunCondition = s.RunCondition
	}

	return out, nil
}

func (s script) execute(event event) {
	if s.Event != event {
		return
	}

	path := s.Path
	if !path.Exists() {
		logger.Warn("not found", "script", path.AbsPath())
		return
	}

	data, err := loadDataConfig()
	if err != nil {
		logger.Warn("unable to load data config", "error", err)
		return
	}

	switch s.RunCondition {
	case runOnce:
		if data.alreadyExecuted(s) {
			logger.Debug("already executed", "script", path.AbsPath())
			return
		}
	case runChanged:
		if !data.hashChanged(s) {
			logger.Debug("hash not changed", "script", path.AbsPath())
			return
		}
	}

	logger.Info("execute", "script", path.AbsPath())
	if err := cmd.Run(path.AbsPath()); err != nil {
		logger.Warn("failed to execute", "script", path.AbsPath(), "error", err)
		return
	}

	logger.Debug("successfully executed", "script", path.AbsPath())

	if s.RunCondition == runOnce || s.RunCondition == runChanged {
		if err := data.addScript(s); err != nil {
			logger.Error("failed to write data config", "error", err)
		}
	}
}
