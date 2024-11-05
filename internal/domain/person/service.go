package person

import (
	"context"

	"github.com/felipeversiane/go-otel/internal/infra/config/rest"
)

type personService struct {
	repository PersonRepositoryInterface
}

type PersonServiceInterface interface {
	InsertOneService(req PersonRequest, ctx context.Context) (string, *rest.RestError)
}

func NewPersonService(repository PersonRepositoryInterface) PersonServiceInterface {
	return &personService{repository}
}

func (service *personService) InsertOneService(req PersonRequest, ctx context.Context) (string, *rest.RestError) {
	domain := ConvertRequestToDomain(req)
	id, err := service.repository.InsertOneRepository(domain, ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}
