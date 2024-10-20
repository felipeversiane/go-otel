package person

type personService struct {
	repository PersonRepositoryInterface
}

type PersonServiceInterface interface {
}

func NewPersonService(repository PersonRepositoryInterface) PersonServiceInterface {
	return &personService{repository}
}
