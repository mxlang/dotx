package config

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/fs"
	"github.com/mxlang/dotx/internal/logger"
)

type executedScript struct {
	Path fs.Path `yaml:"path"`
	Hash string  `yaml:"hash"`
}

func (es *executedScript) UnmarshalYAML(unmarshal func(any) error) error {
	var temp struct {
		Path string `yaml:"path"`
		Hash string `yaml:"hash"`
	}

	if err := unmarshal(&temp); err != nil {
		return err
	}

	es.Path = repoDirPath().Join(temp.Path)
	es.Hash = temp.Hash

	return nil
}

func (es executedScript) MarshalYAML() (any, error) {
	type temp struct {
		Path string `yaml:"path"`
		Hash string `yaml:"hash"`
	}

	out := temp{
		Path: strings.Replace(es.Path.AbsPath(), repoDirPath().AbsPath(), "", 1),
		Hash: es.Hash,
	}

	return out, nil
}

type dataConfig struct {
	Scripts []executedScript `yaml:"scripts"`
}

func (d dataConfig) alreadyExecuted(s script) bool {
	for _, r := range d.Scripts {
		if r.Path == s.Path {
			return true
		}
	}

	return false
}

func (d dataConfig) hashChanged(s script) bool {
	idx, found := d.findRecord(s.Path)
	if !found {
		return true
	}

	currHash, err := computeFileHash(s.Path)
	if err != nil {
		logger.Warn("failed to compute script hash", "script", s.Path.AbsPath(), "error", err)
		return true
	}

	return d.Scripts[idx].Hash != currHash
}

func (d dataConfig) addScript(s script) error {
	h, err := computeFileHash(s.Path)
	if err != nil {
		return fmt.Errorf("unable to compute script hash: %w", err)
	}

	if idx, found := d.findRecord(s.Path); found {
		d.Scripts[idx].Hash = h
	} else {
		d.Scripts = append(d.Scripts, executedScript{Path: s.Path, Hash: h})
	}

	return d.save()
}

func (d dataConfig) findRecord(p fs.Path) (int, bool) {
	for i, r := range d.Scripts {
		if r.Path == p {
			return i, true
		}
	}

	return -1, false
}

func (d dataConfig) save() error {
	b, err := yaml.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data state: %w", err)
	}

	if err := os.WriteFile(dataConfigFilePath().AbsPath(), b, 0644); err != nil {
		return fmt.Errorf("unable to write data state: %w", err)
	}

	return nil
}

func computeFileHash(path fs.Path) (string, error) {
	file, err := os.Open(path.AbsPath())
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func loadDataConfig() (dataConfig, error) {
	content, err := os.ReadFile(dataConfigFilePath().AbsPath())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading data config", "error", err)
		}
	}

	config := dataConfig{}
	if err := yaml.Unmarshal(content, &config); err != nil {
		return config, fmt.Errorf("invalid data config: %w", err)
	}

	return config, nil
}
