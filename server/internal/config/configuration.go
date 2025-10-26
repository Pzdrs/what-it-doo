package config

type Configuration struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DBConfig       `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Gravatar GravatarConfig `mapstructure:"gravatar"`
}

type ServerConfig struct {
	Port int `validate:"min=1,max=65535"`
}

type DBConfig struct {
	URL      string
	Host     string
	Port     int `validate:"min=1,max=65535"`
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host     string
	Port     int `validate:"min=1,max=65535"`
	Password string
}

type GravatarConfig struct {
	Enabled bool
	Url     string
}
