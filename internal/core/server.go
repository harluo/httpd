package core

import (
	"net/http"

	"github.com/harluo/httpd/internal/config"
)

type Server struct {
	http   *http.Server
	config *config.Server
}

func newServer(config *config.Server) (server *Server) {
	server = new(Server)
	server.http = new(http.Server)
	server.http.Addr = config.Addr()

	if nil != config.Timeout && 0 != config.Timeout.Read {
		server.http.WriteTimeout = config.Timeout.Read
	}
	if nil != config.Timeout && 0 != config.Timeout.Write {
		server.http.WriteTimeout = config.Timeout.Write
	}

	return
}

func (s *Server) Http() *http.Server {
	return s.http
}
func (s *Server) Port() uint16 {
	return s.config.Port
}
