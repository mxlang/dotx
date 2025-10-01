package config

import (
	"fmt"
	"path/filepath"

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
	Path         string       `yaml:"path"` // TODO change type to fs.Path
	Event        event        `yaml:"on"`
	RunCondition runCondition `yaml:"run,omitempty"`
}

func (s *script) execute(event event) {
	if s.Event != event {
		return
	}

	path := fs.NewPath(filepath.Join(repoDirPath(), s.Path))
	if !path.Exists() {
		logger.Warn("not found", "script", path.AbsPath())
		return
	}

	data := loadDataConfig()

	switch s.RunCondition {
	case runOnce:
		if data.alreadyExecuted(*s) {
			logger.Debug("already executed", "script", path.AbsPath())
			return
		}

		if err := data.addScript(*s); err != nil {
			logger.Error("failed to write data config", "error", err)
		}
	case runChanged:
		if data.alreadyExecuted(*s) {
			if data.hashChanged(*s) {
				logger.Debug("hash changed", "script", path.AbsPath())
				if err := data.updateScript(*s); err != nil {
					logger.Error("failed to write data config", "error", err)
				}
			} else {
				logger.Debug("hash not changed", "script", path.AbsPath())
				return
			}
		} else {
			if err := data.addScript(*s); err != nil {
				logger.Error("failed to write data config", "error", err)
			}
		}
	}

	logger.Info("execute", "script", path.AbsPath(), "on", s.Event, "run", s.RunCondition)
	if err := cmd.Run(path.AbsPath()); err != nil {
		logger.Warn("failed to execute", "script", path.AbsPath(), "error", err)
	} else {
		logger.Debug("successfully executed", "script", path.AbsPath())
	}
}
