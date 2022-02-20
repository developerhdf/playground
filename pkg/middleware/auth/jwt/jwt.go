package jwt

import (
	"net/http"

	"example/richard/sovtech/pkg/repositories"
	"example/richard/sovtech/pkg/util/auth"
	jwtutil "example/richard/sovtech/pkg/util/auth/jwt"
)

//does jwt token auth check, falls back to basic if unsuccessful

type JWTAuth struct {
	userRepository *repositories.UserRepository
}

func NewJWTAuth(ur *repositories.UserRepository) *JWTAuth {
	return &JWTAuth{ur}
}

func (ja JWTAuth) getJWTUser(authHeader string) *models.User {
	var user *models.User
	if username, jwtErr := jwtutil.GetTokenUsername(authHeader); err == nil {
		if repoUser, repoErr := ja.userRepository.GetUser(username); repoErr == nil {
			user = repoUser
		}
	}
	return user
}

func (ja JWTAuth) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		user := ja.getJWTUser(header)
		if user == nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//add user to context
		auth.AddUserToRequestContext(r, user)
		next.ServeHTTP(w, r)
	})
}
