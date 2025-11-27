package core

import (
	"net/http"

	"github.com/harluo/httpd/internal/config"
)

type Server struct {
	*http.Server

	config *config.Server
}

func (s *Server) Port() uint16 {
	return s.config.Port
}
