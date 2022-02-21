package auth

import (
	"net/http"
)

type AuthMiddleware interface {
	Authenticate(handler http.Handler) http.Handler
}
