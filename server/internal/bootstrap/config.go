package bootstrap

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"pycrs.cz/what-it-doo/internal/config"
)

func InitConfig() (config.Configuration, error) {
	v := newViper()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config.Configuration{}, fmt.Errorf("failed to read config: %w", err)
		}
		log.Println("No configuration file found")
	} else {
		log.Println("Using configuration file:", v.ConfigFileUsed())
	}

	var cfg config.Configuration
	if err := v.Unmarshal(&cfg); err != nil {
		return config.Configuration{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterStructValidation(
		config.DbConfigStructLevelValidation,
		config.DBConfig{},
	)
	if err := validate.Struct(cfg); err != nil {
		return config.Configuration{}, fmt.Errorf("invalid config: %w", err)
	}

	fmt.Println(cfg)

	return cfg, nil
}

func newViper() *viper.Viper {
	v := viper.NewWithOptions(viper.ExperimentalBindStruct())

	v.SetConfigName("wid")
	v.SetEnvPrefix("WID")

	v.AddConfigPath(".")
	v.AddConfigPath("/etc/whatitdoo")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return v
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("server.port", 8080)
	v.SetDefault("database.port", 5432)
	v.SetDefault("redis.port", 6379)

	v.SetDefault("gravatar.enabled", true)
	v.SetDefault("gravatar.url", "https://secure.gravatar.com/avatar/{{hash}}?s={{size}}&d=identicon")
}
