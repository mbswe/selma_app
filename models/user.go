package models

import (
	"github.com/mbswe/selma/model"
)

// User represents a user in the system
type User struct {
	model.ModelBase
	Name  string
	Email string
}

// NewUser creates a new User instance
func NewUser(id int, name, email string) *User {
	return &User{
		ModelBase: model.ModelBase{ID: id},
		Name:      name,
		Email:     email,
	}
}
