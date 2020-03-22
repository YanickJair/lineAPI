package models

import (
	"apiv1/src/db"

	"github.com/google/uuid"
)

//User - user model
type User struct {
	id        string
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string `json:"last_name" xml:"last_name" form:"last_name"`
	Avatar    string `json:"avatar" xml:"avatar" form:"avatar"`
}

//CreateAccount - register a new user in the system
func (user *User) CreateAccount() {
	db := db.Connector()

	defer db.Close()

	ID, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}
	user.id = ID.String()

	err = db.Insert(user)
	if err != nil {
		panic(err)
	}
}
