package person

import (
	"time"

	"github.com/felipeversiane/go-otel/internal/domain"
)

type PersonRequest struct {
	Nickname string   `json:"nickname" binding:"required,min=2,max=20"`
	Name     string   `json:"name" binding:"required,min=2,max=50"`
	Stack    []string `json:"stack"`
}

type PersonResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Name      string    `json:"name"`
	Stack     []string  `json:"stack"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertRequestToDomain(req PersonRequest) *domain.Person {
	return domain.NewPerson(req.Nickname, req.Name, req.Stack)
}

func ConvertDomainToResponse(domain *domain.Person) *PersonResponse {
	return &PersonResponse{domain.ID, domain.Nickname, domain.Name, domain.Stack, domain.CreatedAt}
}
