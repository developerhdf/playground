package graph

import (
	"crypto/tls"
	"net/http"

	"example/richard/sovtech/pkg/repositories"
	"example/richard/sovtech/pkg/repositories/memory"
	"example/richard/sovtech/pkg/repositories/swapi"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PeopleRepo repositories.PeopleRepository
	UserRepo   repositories.UserRepository
}

func NewSovTechResolver() *Resolver {
	//at the time of writing this, the swapi certificate had expired
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	return &Resolver{
		PeopleRepo: swapi.NewPeopleRepository(httpClient),
		UserRepo:   memory.NewUserRepository(),
	}
}
