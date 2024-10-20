package domain

import (
	"strings"

	"github.com/google/uuid"
)

type Person struct {
	ID        string
	Nickname  string
	Name      string
	Birthdate string
	Stack     []string
}

func NewPerson(
	nickname string,
	name string,
	birthdate string,
	stack []string,
) *Person {
	return &Person{
		ID:        uuid.NewString(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     stack,
	}
}

func (p *Person) StackStr() string {
	return strings.Join(p.Stack, ",")
}
