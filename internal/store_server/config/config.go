package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var c Config

type Config struct {
	Platform PlatformConfig `mapstructure:"platform"`
	DB       DatabaseConfig `mapstructure:"db"`
	Auth     AuthConfig     `mapstructure:"auth"`
}

type AuthConfig struct {
	SecretKey   string        `mapstructure:"secret_key"`   // a 16-32 character random string for auth
	TokenExpiry time.Duration `mapstructure:"token_expiry"` // how long the auth token should be valid in days
}

const Delimiter = "::"

type PlatformConfig struct {
	Env     string `mapstructure:"env"`
	Name    string `mapstructure:"name"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
	Host                string `mapstructure:"host"`
	Port                int    `mapstructure:"port"`
	User                string `mapstructure:"user"`
	Password            string `mapstructure:"password"`
	Schema              string `mapstructure:"schema"`
	MigrateSchema       bool   `mapstructure:"migrate_schema"`
	MigrateData         bool   `mapstructure:"migrate_data"`
	MaxIdealConnections int    `mapstructure:"max_ideal_connections"`
	Debug               bool   `mapstructure:"debug"`
}

func Load(path string) {
	configName := os.Getenv("CONFIG")
	if configName == "" {
		configName = "config"
	}
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)       // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("can not load config: %s", err)
	}
}

func Get() Config {
	return c
}
