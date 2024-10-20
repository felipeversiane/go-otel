package person

type personController struct {
	service PersonServiceInterface
}

type PersonControllerInterface interface {
}

func NewPersonController(service PersonServiceInterface) PersonControllerInterface {
	return &personController{service}
}
