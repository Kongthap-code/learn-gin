package model

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func NewUser(name string) *User {
	return &User{
		Name:  name,
	}
}
