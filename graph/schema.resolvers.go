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
	jwtToken := ""
	var err error
	modelsUser := ToModelsUser(input)
	if createUserErr := r.UserRepo.Create(modelsUser); createUserErr == nil {
		if token, tokenErr := GetJWTToken(modelsUser); tokenErr == nil {
			jwtToken = token
		} else {
			err = tokenErr
		}
	} else {
		err = createUserErr
	}
	return jwtToken, err
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	jwtToken := ""
	err := fmt.Errorf("incorrect email or password")
	if repoUser, repoErr := r.UserRepo.GetUser(input.Email); repoErr == nil {
		if ValidPassword(input.Password, repoUser.Password) {
			if token, tokenErr := GetJWTToken(repoUser); tokenErr == nil {
				jwtToken = token
				err = nil
			} else {
				err = tokenErr
			}
		}
	}
	return jwtToken, err
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	jwtToken := ""
	var err error
	if username, tokenErr := GetTokenUsername(input.Token); tokenErr == nil {
		if newToken, creationErr := CreateToken(username); creationErr == nil {
			jwtToken = newToken
		} else {
			err = fmt.Errorf("could not create new token")
		}
	} else {
		err = fmt.Errorf("invalid token")
	}
	return jwtToken, err
}

func (r *peopleResultResolver) People(ctx context.Context, obj *model.PeopleResult) ([]*model.Person, error) {
	people := make([]*model.Person, 0)
	switch {
	case GetContextUser(ctx) == nil:
		return people, fmt.Errorf("Access denied")
	case obj == nil:
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
	if GetContextUser(ctx) == nil {
		return people, fmt.Errorf("Access denied")
	}
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
	if GetContextUser(ctx) == nil {
		return person, fmt.Errorf("Access denied")
	}
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
