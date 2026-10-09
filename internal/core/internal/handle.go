package internal

import (
	"net/http"
)

type Handle struct {
	Path    string
	Method  string
	Handler http.Handler
}
