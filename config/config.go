package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		PG    `yaml:"postgres"`
		Log   `yaml:"logger"`
		Auth  `yaml:"auth"`
		Redis `yaml:"redis"`
	}

	App struct {
		Name    string `env:"APP_NAME" yaml:"name"`
		Version string `yaml:"version" env:"APP_VERSION"`
		IsDebug bool   `yaml:"is_debug" env:"APP_IS_DEBUG"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}

	Auth struct {
		TTL string `yaml:"ttl" env:"AUTH_TTL"` //TODD FIX move env variable
	}

	Redis struct {
		Host       string `yaml:"host" env:"REDIS_HOST"`
		Password   string `yaml:"password" env:"REDIS_PASSWORD"`
		Username   string `yaml:"username" env:"REDIS_USERNAME"`
		Port       string `yaml:"port" env:"REDIS_PORT"`
		DB         int    `yaml:"db_name" env:"REDIS_DB"`
		ExpireTime string `yaml:"expire_time" env:"REDIS_EXPIRE_TIME"`
	}

	PG struct {
		Host                string `yaml:"host" env:"POSTGRES_HOST"`
		DbName              string `yaml:"db_name" env:"POSTGRES_DB"`
		Username            string `yaml:"username" env:"POSTGRES_USERNAME"`
		Password            string `yaml:"password" env:"POSTGRES_PASSWORD"`
		Port                string `yaml:"port" env:"POSTGRES_PORT"`
		AttemptToConnect    int    `yaml:"attempt_to_connect" env:"POSTGRES_ATTEMPTS"`
		MigrationsSourceURL string `yaml:"migrations_source_url" env:"POSTGRES_MIGRATIONS_SOURCE_URL"`
	}
)

func NewConfig() (*Config, error) { //TODO add logger
	cfg := &Config{}
	//once.Do(func() {
	//	if err := cleanenv.ReadConfig("config/.env", cfg); err != nil {
	//		return cfg, err
	//	}
	//
	//	if err := cleanenv.ReadEnv(cfg); err != nil {
	//		return cfg, err
	//	}
	//})

	if err := cleanenv.ReadConfig(".env", cfg); err != nil {
		return cfg, err
	}

	if err := cleanenv.ReadConfig("config/config.yml", cfg); err != nil {
		return cfg, err
	}

	//if err := cleanenv.ReadEnv(cfg); err != nil {
	//	return cfg, err
	//}

	return cfg, nil

}
