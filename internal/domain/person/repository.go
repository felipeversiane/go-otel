package person

import (
	"context"
	"fmt"
	"strings"

	"github.com/felipeversiane/go-otel/internal/domain"
	"github.com/felipeversiane/go-otel/internal/infra/config/rest"
	"github.com/felipeversiane/go-otel/internal/infra/services/database"
)

type personRepository struct {
	db database.DatabaseInterface
}

type PersonRepositoryInterface interface {
	InsertOneRepository(domain *domain.Person, ctx context.Context) (string, *rest.RestError)
}

func NewPersonRepository(db database.DatabaseInterface) PersonRepositoryInterface {
	return &personRepository{db}
}

func (repository *personRepository) InsertOneRepository(domain *domain.Person, ctx context.Context) (string, *rest.RestError) {
	search := fmt.Sprintf("%s %s %s", domain.Nickname, domain.Name, strings.Join(domain.Stack, " "))
	query := `INSERT INTO persons (nickname, name, birthdate, stack, search) VALUES (:nickname, :name, :birthdate, :stack, :search) RETURNING id`
	args := map[string]interface{}{
		"nickname":  domain.Nickname,
		"name":      domain.Name,
		"birthdate": domain.Birthdate,
		"stack":     domain.Stack,
		"search":    search,
	}
	var id string
	err := repository.db.GetDB().QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		return "", rest.NewInternalServerError(fmt.Sprintf("unable to insert person: %v", err))
	}
	return id, nil
}
