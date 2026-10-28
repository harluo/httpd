package internal

import "net/http"

type Health struct {
	// 无字段
}

func NewHealth() *Health {
	return new(Health)
}

func (h *Health) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("ok\n"))
}