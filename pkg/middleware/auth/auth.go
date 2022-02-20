package auth

type AuthMiddleware interface {
	Authenticate(handler http.HandlerFunc) http.HandlerFunc
}
