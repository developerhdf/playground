package main

import (
	"context"
	"fmt"
	"net/http"

	"example/richard/sovtech/repository"
	"example/richard/sovtech/repository/swapi"
)

func printData(repo repository.PeopleRepository) {
	result, err := repo.GetPeople(context.Background(), 1)
	fmt.Printf("%v %s\n", result, err)
	result, err = repo.SearchPeople(context.Background(), "luke", 1)
	fmt.Printf("%v %s\n", result, err)
	result, err = repo.SearchPeople(context.Background(), "darth", 1)
	fmt.Printf("%v %s\n", result, err)
}

func main() {
	repo := swapi.NewPeopleRepository(&http.Client{})
	printData(repo)
}
