package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env  string `mapstructure:"APP_ENV"`
	Port string `mapstructure:"PORT"`

	DBConfig    DBConfig    `mapstructure:",squash"`
	CacheConfig CacheConfig `mapstructure:",squash"`
}

type DBConfig struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_DATABASE"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

type CacheConfig struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisUsername string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func Load(logger *slog.Logger) (*Config, error) {
	v := viper.New() // Create a new Viper instance to avoid any global state issues

	// Configure Viper
	v.SetEnvPrefix("") // No prefix for env vars
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set all environment variables into viper explicitly
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		v.Set(pair[0], pair[1])
	}

	// Create config struct
	var config Config

	// Unmarshal into struct
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
