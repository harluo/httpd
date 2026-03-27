package config

import (
	"time"
)

type Timeout struct {
	// 读数据
	Read time.Duration `json:"read,omitempty"`
	// 写数据
	Write time.Duration `json:"write,omitempty"`
	// 空闲
	Idle time.Duration `json:"idle,omitempty"`
}

func newTimeout(server *Server) *Timeout {
	return server.Timeout
}
