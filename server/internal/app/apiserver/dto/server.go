package dto

type ServerInfo struct {
	Version string `json:"version" validate:"required"`
}

type ServerConfig struct {
}
