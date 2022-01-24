package config

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
		Log  `yaml:"logger"`
	}

	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
		IsDebug bool   `yaml:"is_debug" env:"APP_IS_DEBUG"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}

	PG struct {
		DbName   string `yaml:"db_name" env:"PG_DATABASE_NAME"`
		Username string `yaml:"username" env:"PG_USERNAME"`
		Password string `yaml:"password" env:"PG_PASSWORD"`
		Port     string `yaml:"port" env:"PG_PORT"`
	}
)
