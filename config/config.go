package config

type Config struct {
	appConfig
	repoConfig
}

func Load() Config {
	appConfig := loadAppConfig()
	return Config{
		appConfig:  appConfig,
		repoConfig: loadRepoConfig(appConfig),
	}
}
