package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   `mapstructure:"app"`
	Http  `mapstructure:"http"`
	Log   `mapstructure:"log"`
	MySQL `mapstructure:"mysql"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Http struct {
	Port string `mapstructure:"port"`
}

type Log struct {
	Level string `mapstructure:"level"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

// LoadConfig reads the configuration file.
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
