package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	TgAPIToken    string `yaml:"tg_api_token"`
	WeatherAPIKey string `yaml:"weather_api_key"`
	LogLevel      string `yaml:"log_level"`
}

func MustReadConfig(configPath string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
