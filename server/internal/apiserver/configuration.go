package apiserver

type Configuration struct {
	Server   serverConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database" validate:"required"`
}

type serverConfiguration struct {
	Port int `validate:"min=1,max=65535"`
}

type DatabaseConfiguration struct {
	URL      string `mapstructure:"url"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}
