package config

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/mxlang/dotx/internal/logger"
)

type executed struct {
	Path string `yaml:"path"` // TODO change type to fs.Path
	Hash string `yaml:"hash"`
}

type dataConfig struct {
	Scripts []executed `yaml:"scripts"`
}

func (d *dataConfig) alreadyExecuted(s script) bool {
	for _, exec := range d.Scripts {
		if exec.Path == s.Path {
			return true
		}
	}

	return false
}

func (d *dataConfig) hashChanged(s script) bool {
	for _, exec := range d.Scripts {
		if exec.Path == s.Path {
			hash, err := getHash(s)
			if err != nil {
				// TODO
			}

			if exec.Hash != hash {
				return true
			}
		}
	}

	return false
}

func (d *dataConfig) addScript(s script) error {
	hash, err := getHash(s)
	if err != nil {
		return fmt.Errorf("unable to create hash: %w", err)
	}

	exec := executed{
		Path: s.Path,
		Hash: hash,
	}

	d.Scripts = append(d.Scripts, exec)

	config, err := yaml.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data config: %w", err)
	}

	if err := os.WriteFile(dataConfigFilePath(), config, 0644); err != nil {
		return fmt.Errorf("unable to write data config: %w", err)
	}

	return nil
}

func (d *dataConfig) updateScript(s script) error {
	hash, err := getHash(s)
	if err != nil {
		return fmt.Errorf("unable to create hash: %w", err)
	}

	for i, exec := range d.Scripts {
		if exec.Path == s.Path {
			d.Scripts[i].Hash = hash
		}
	}

	config, err := yaml.Marshal(d)
	if err != nil {
		return fmt.Errorf("unable to marshal data config: %w", err)
	}

	if err := os.WriteFile(dataConfigFilePath(), config, 0644); err != nil {
		return fmt.Errorf("unable to write data config: %w", err)
	}

	return nil
}

func getHash(s script) (string, error) {
	file, err := os.Open(filepath.Join(repoDirPath(), s.Path))
	if err != nil {
		return "", err
	}
	defer file.Close()

	sha := sha256.New()
	if _, err := io.Copy(sha, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha.Sum(nil)), nil
}

func loadDataConfig() dataConfig {
	content, err := os.ReadFile(dataConfigFilePath())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warn("error while reading data config", "error", err)
		}
	}

	config := dataConfig{}

	if err := yaml.Unmarshal(content, &config); err != nil {
		logger.Error("invalid data config", "error", err)
	}

	return config
}
