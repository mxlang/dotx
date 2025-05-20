package config

type appConfig struct {
	Verbose       bool   `yaml:"verbose"`
	CommitMessage string `yaml:"commitMessage"`
	DeployOnInit  bool   `yaml:"deployOnInit"`
	DeployOnPull  bool   `yaml:"deployOnPull"`
}
