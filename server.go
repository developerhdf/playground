package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"example/richard/sovtech/pkg/middleware/auth/jwt"
	"example/richard/sovtech/pkg/repositories/memory"
	"example/richard/sovtech/pkg/repositories/swapi"

	"example/richard/sovtech/graph"
	"example/richard/sovtech/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//at the time of writing this, the swapi certificate had expired
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	peopleRepo := swapi.NewPeopleRepository(httpClient)
	userRepo := memory.NewUserRepository()
	jwtAuth := jwt.NewJWTAuth(userRepo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewSovTechResolver(
		userRepo,
		peopleRepo,
	)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", jwtAuth.Authenticate(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
