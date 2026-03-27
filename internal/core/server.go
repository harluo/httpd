package core

import (
	"crypto/tls"
	"net/http"

	"github.com/harluo/httpd/internal/config"
)

type Server struct {
	http   *http.Server
	config *config.Server
}

func newServer(config *config.Server) (server *Server) {
	server = new(Server)
	server.config = config

	server.http = new(http.Server)
	server.http.Addr = config.Addr()
	server.http.TLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13, // 启用高版本
		NextProtos: []string{
			"h2",
			"http/1.1",
		},

		// 优化证书配置
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},

		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP256,
		},

		// 会话票证，提高性能
		SessionTicketsDisabled: false,
	}

	if config.Timeout != nil && config.Timeout.Read != 0 {
		server.http.WriteTimeout = config.Timeout.Read
	}
	if config.Timeout != nil && config.Timeout.Write != 0 {
		server.http.WriteTimeout = config.Timeout.Write
	}
	if config.Timeout != nil && config.Timeout.Idle != 0 {
		server.http.IdleTimeout = config.Timeout.Idle
	}

	return
}

func (s *Server) Http() *http.Server {
	return s.http
}

func (s *Server) Port() uint16 {
	return s.config.Port
}

func (s *Server) Addr() string {
	return s.config.Addr()
}
