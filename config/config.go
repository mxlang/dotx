package config

type Config struct {
	appConfig
	repoConfig
}

func Load() Config {
	return Config{
		appConfig:  loadAppConfig(),
		repoConfig: loadRepoConfig(),
	}
}
