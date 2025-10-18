package apiserver

type Configuration struct {
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
}

type ServerConfiguration struct {
	Port int `validate:"min=1,max=65535"`
}

type DatabaseConfiguration struct {
	URL      string
	Host     string
	Port     int `validate:"min=1,max=65535"`
	User     string
	Password string
	Name     string
}
