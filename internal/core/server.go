package core

import (
	"net/http"
	"path/filepath"

	"github.com/harluo/httpd/internal/config"
	"github.com/harluo/httpd/internal/core/internal"
)

type Server struct {
	http    *http.Server
	config  *config.Server
	handles []*internal.Handle
}

func newServer(config *config.Server) (server *Server) {
	server = new(Server)
	server.config = config

	server.http = new(http.Server)
	server.http.Addr = config.Addr()

	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	server.http.Protocols = protocols

	if config.Timeout != nil && config.Timeout.Header != 0 {
		server.http.ReadHeaderTimeout = config.Timeout.Header
		server.http.ReadTimeout = config.Timeout.Read
		server.http.WriteTimeout = config.Timeout.Write
		server.http.IdleTimeout = config.Timeout.Idle
	}

	health := internal.NewHealth()
	server.handles = append(server.handles, &internal.Handle{ // 默认健康检查
		Path:    "/health",
		Method:  "GET",
		Handler: health,
	}, &internal.Handle{ // 默认准备完成探针
		Path:    "/ready",
		Method:  "GET",
		Handler: health,
	})

	return
}

func (s *Server) Http() *http.Server {
	return s.http
}

func (s *Server) Basepath() string {
	return s.config.Basepath
}

func (s *Server) Path(required string, optionals ...string) string {
	paths := make([]string, 0, len(optionals)+2)
	paths = append(paths, s.Basepath())
	paths = append(paths, required)
	paths = append(paths, optionals...)

	return filepath.ToSlash(filepath.Join(paths...))
}

func (s *Server) Port() uint16 {
	return s.config.Port
}

func (s *Server) Addr() string {
	return s.config.Addr()
}

func (s *Server) Handles() []*internal.Handle {
	return s.handles
}

func (s *Server) Get(path string, handler http.Handler) (server *Server) {
	s.handles = append(s.handles, &internal.Handle{
		Path:    path,
		Method:  http.MethodGet,
		Handler: handler,
	})

	return
}

func (s *Server) Put(path string, handler http.Handler) (server *Server) {
	s.handles = append(s.handles, &internal.Handle{
		Path:    path,
		Method:  http.MethodPut,
		Handler: handler,
	})

	return
}

func (s *Server) Post(path string, handler http.Handler) (server *Server) {
	s.handles = append(s.handles, &internal.Handle{
		Path:    path,
		Method:  http.MethodPost,
		Handler: handler,
	})

	return
}

func (s *Server) Delete(path string, handler http.Handler) (server *Server) {
	s.handles = append(s.handles, &internal.Handle{
		Path:    path,
		Method:  http.MethodDelete,
		Handler: handler,
	})

	return
}

func (s *Server) Patch(path string, handler http.Handler) (server *Server) {
	s.handles = append(s.handles, &internal.Handle{
		Path:    path,
		Method:  http.MethodPatch,
		Handler: handler,
	})

	return
}
