package controller

import (
	"net/http"

	"pycrs.cz/what-it-doo/internal/app/apiserver/common"
	"pycrs.cz/what-it-doo/internal/app/apiserver/dto"
	"pycrs.cz/what-it-doo/pkg/version"
)

type ServerController struct {
}

func NewServerController() *ServerController {
	return &ServerController{}
}

// HandleAbout retrieves the server information
//
//	@Summary		Get server information
//	@Description	Get information about the server
//	@Id				getServerInfo
//	@Tags			miscellaneous
//	@Produce		json
//	@Success		200	{object}	dto.ServerInfo
//	@Router			/server/about [get]
func (c *ServerController) HandleAbout(w http.ResponseWriter, r *http.Request) {
	common.Encode(w, r, 200, dto.ServerInfo{
		Version: version.Version,
	})
}

// HandleConfig retrieves the server configuration
//
//	@Summary		Get server configuration
//	@Description	Get server configuration
//	@Id				getServerConfig
//	@Tags			miscellaneous
//	@Produce		json
//	@Success		200	{object}	dto.ServerConfig
//	@Router			/server/config [get]
func (c *ServerController) HandleConfig(w http.ResponseWriter, r *http.Request) {
	common.Encode(w, r, 200, dto.ServerConfig{})
}
