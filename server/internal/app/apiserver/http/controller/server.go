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

// HandleAbout
//
//	@Summary		Get server information
//	@Description	Get information about the server
//	@Id				getServerInfo
//	@Tags			Server
//	@Produce		json
//	@Success		200	{object}	dto.ServerInfo
//	@Router			/server/about [get]
func (c *ServerController) HandleAbout(w http.ResponseWriter, r *http.Request) {
	common.WriteJSON(w, 200, dto.ServerInfo{
		Version: version.Version,
	})
}

// HandleConfig
//
//	@Summary		Get server configuration
//	@Description	Get server configuration
//	@Id				getServerConfig
//	@Tags			Server
//	@Produce		json
//	@Success		200	{object}	dto.ServerConfig
//	@Router			/server/config [get]
func (c *ServerController) HandleConfig(w http.ResponseWriter, r *http.Request) {
	common.WriteJSON(w, 200, dto.ServerConfig{})
}
