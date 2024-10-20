package person

import (
	"github.com/felipeversiane/go-otel/internal/infra/services/database"
)

type personRepository struct {
	db database.DatabaseInterface
}

type PersonRepositoryInterface interface {
}

func NewPersonRepository(db database.DatabaseInterface) PersonRepositoryInterface {
	return &personRepository{db}
}
