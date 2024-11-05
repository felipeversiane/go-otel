package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        string
	Nickname  string
	Name      string
	Stack     []string
	CreatedAt time.Time
}

func NewPerson(
	nickname string,
	name string,
	stack []string,
) *Person {
	return &Person{
		ID:        uuid.NewString(),
		Nickname:  nickname,
		Name:      name,
		Stack:     stack,
		CreatedAt: time.Now(),
	}
}

func (p *Person) StackStr() string {
	return strings.Join(p.Stack, ",")
}
