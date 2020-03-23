package models

import (
	"apiv1/src/db"
	"fmt"

	"github.com/google/uuid"
)

//User - user model
type User struct {
	id        string
	Username  string `json:"username" xml:"username" form:"username"`
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string `json:"last_name" xml:"last_name" form:"last_name"`
	Avatar    string `json:"avatar" xml:"avatar" form:"avatar"`
}

//CreateAccount - register a new user in the system
func (user *User) CreateAccount() {
	db := db.Connector()

	ID, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}
	user.id = ID.String()

	fmt.Println(user)

	err = db.Insert(user)
	if err != nil {
		panic(err)
	}
}

//SignIn - login to the app
func (user *User) SignIn() *User {

	return &User{}
}
