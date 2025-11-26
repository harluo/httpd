package core

import (
	"github.com/harluo/httpd/internal/config"
)

func newHttpServer(config *config.Server) (server *Server) {
	server = new(Server)
	server.Addr = config.Addr()
	if nil != config.Timeout && 0 != config.Timeout.Read {
		server.WriteTimeout = config.Timeout.Read
	}
	if nil != config.Timeout && 0 != config.Timeout.Write {
		server.WriteTimeout = config.Timeout.Write
	}

	return
}
