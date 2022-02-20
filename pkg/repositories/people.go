package repositories

import (
	"context"

	"example/richard/sovtech/pkg/models"
)

type PeopleRepository interface {
	GetPeople(ctx context.Context, page int) (*models.PeopleResult, error)
	SearchPeople(ctx context.Context, name string, page int) (*models.PeopleResult, error)
}
