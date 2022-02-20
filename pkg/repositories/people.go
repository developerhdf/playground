package repositories

import (
	"context"
	"encoding/json"

	"example/richard/sovtech/pkg/models"
)

type PeopleResult struct {
	People  []*models.Person `json:"people"`
	HasMore bool             `json:"hasMore"`
}

func (p PeopleResult) String() string {
	resultString := "{}"
	if jsonResult, err := json.Marshal(p); err == nil {
		resultString = string(jsonResult)
	}
	return resultString
}

type PeopleRepository interface {
	GetPeople(ctx context.Context, page int) (PeopleResult, error)
	SearchPeople(ctx context.Context, name string, page int) (PeopleResult, error)
}
