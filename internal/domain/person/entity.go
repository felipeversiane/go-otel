package person

import "github.com/felipeversiane/go-otel/internal/domain"

type PersonRequest struct {
	Nickname  string   `json:"nickname" binding:"required,unique,min=2,max=20"`
	Name      string   `json:"name" binding:"required,min=2,max=50"`
	Birthdate string   `json:"birthdate" binding:"required,datetime"`
	Stack     []string `json:"stack"`
}

type PersonResponse struct {
	ID        string   `json:"id"`
	Nickname  string   `json:"nickname"`
	Name      string   `json:"name"`
	Birthdate string   `json:"birthdate"`
	Stack     []string `json:"stack"`
}

func ConvertRequestToDomain(req *PersonRequest) *domain.Person {
	return domain.NewPerson(req.Nickname, req.Name, req.Birthdate, req.Stack)
}

func ConvertDomainToResponse(domain *domain.Person) *PersonResponse {
	return &PersonResponse{domain.ID, domain.Nickname, domain.Name, domain.Birthdate, domain.Stack}
}
