package config

import (
	"bytes"
	_ "embed"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var defaultConfiguration []byte

var CNF *Config

type Postgres struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type App struct {
	UrlBase     string `mapstructure:"url-base"`
	UrlReset    string `mapstructure:"url-reset"`
	UrlActivate string `mapstructure:"url-activate"`
	Port        string `mapstructure:"port"`
}

type Jwt struct {
	TokenTTL   int    `mapstructure:"token-ttl"`
	PrivateKey string `mapstructure:"private-key"`
}

type Config struct {
	Postgres *Postgres
	App      *App
	Jwt      *Jwt
}

func Read() error {
	viper.SetConfigName("default") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read value ENV variable

	// Read configuration
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return err
	}

	if err := viper.Unmarshal(&CNF); err != nil {
		return err
	}
	return nil
}
