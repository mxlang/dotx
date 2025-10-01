package config

type Config struct {
	RepoPath string // TODO change type to fs.Path

	App  appConfig
	Repo repoConfig
}

func Load() *Config {
	app := loadAppConfig()
	repo := loadRepoConfig()

	return &Config{
		RepoPath: repoDirPath(),

		App:  app,
		Repo: repo,
	}
}
