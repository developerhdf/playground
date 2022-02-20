package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/richard/sovtech/graph/generated"
	"example/richard/sovtech/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *peopleResultResolver) People(ctx context.Context, obj *model.PeopleResult) ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	if obj == nil {
		return people, fmt.Errorf("no results")
	}
	for _, person := range obj.People {
		gqlPerson := model.Person(*person)
		people = append(people, &gqlPerson)
	}
	return people, nil
}

func (r *queryResolver) GetPeople(ctx context.Context, page int) (*model.PeopleResult, error) {
	var err error
	var people *model.PeopleResult
	if repoPeople, repoErr := r.PeopleRepo.GetPeople(ctx, page); repoErr == nil && repoPeople != nil {
		peopleView := model.PeopleResult(*repoPeople)
		people = &peopleView
	} else {
		err = fmt.Errorf("failed to retrieve people")
	}
	return people, err
}

func (r *queryResolver) SearchPeople(ctx context.Context, name string) (*model.Person, error) {
	var err error
	var person *model.Person
	if repoPeople, repoErr := r.PeopleRepo.SearchPeople(ctx, name, 1); repoErr == nil && repoPeople != nil {
		if len(repoPeople.People) > 0 {
			repoPerson := repoPeople.People[0]
			personView := model.Person(*repoPerson)
			person = &personView
		} else {
			err = fmt.Errorf("no results")
		}
	} else {
		err = fmt.Errorf("failed to search people")
	}
	return person, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PeopleResult returns generated.PeopleResultResolver implementation.
func (r *Resolver) PeopleResult() generated.PeopleResultResolver { return &peopleResultResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type peopleResultResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
