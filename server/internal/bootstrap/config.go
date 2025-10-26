package bootstrap

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"pycrs.cz/what-it-doo/internal/apiserver"
	"pycrs.cz/what-it-doo/internal/validation"
)

func InitConfig() (apiserver.Configuration, error) {
	config := newViper()

	setDefaults(config)

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return apiserver.Configuration{}, fmt.Errorf("failed to read config: %w", err)
		}
		log.Println("No configuration file found")
	} else {
		log.Println("Using configuration file:", config.ConfigFileUsed())
	}

	var cfg apiserver.Configuration
	if err := config.Unmarshal(&cfg); err != nil {
		return apiserver.Configuration{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterStructValidation(
		validation.DbConfigStructLevelValidation,
		apiserver.DatabaseConfiguration{},
	)
	if err := validate.Struct(cfg); err != nil {
		return apiserver.Configuration{}, fmt.Errorf("invalid config: %w", err)
	}

	return cfg, nil
}

func newViper() *viper.Viper {
	config := viper.NewWithOptions(viper.ExperimentalBindStruct())

	config.SetConfigName("wid")
	config.SetEnvPrefix("WID")

	config.AddConfigPath(".")
	config.AddConfigPath("/etc/whatitdoo")

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return config
}

func setDefaults(config *viper.Viper) {
	config.SetDefault("server.port", 8080)
	config.SetDefault("database.port", 5432)
}
