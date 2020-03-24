package models

import (
	"apiv1/src/db"
	"apiv1/src/errors"
	"apiv1/src/utils"

	"github.com/google/uuid"
)

//User - user model
type User struct {
	ID        string
	FirstName string `json:"first_name" xml:"first_name" form:"first_name"`
	LastName  string `json:"last_name" xml:"last_name" form:"last_name"`
	Avatar    string `json:"avatar" xml:"avatar" form:"avatar"`
	Password  string `json:"password" xml:"password" form:"password"`
	Email     string `json:"email" xml:"email" form:"email"`
}

//CreateAccount - register a new user in the system
func (user *User) CreateAccount() errors.ErrorCode {
	db := db.Connector()

	ID, err := uuid.NewUUID()

	if err != nil {
		panic(err)
	}

	user.ID = ID.String()
	password, err := utils.HashPassword(user.Password)

	if err != nil {
		panic(err)
	}
	user.Password = password

	err = db.Insert(user)
	if err != nil {
		panic(err)
	}
	return errors.ErrorCode{Message: "Account created", Code: 200}
}

//SignIn - login to the app
func SignIn(email, password string) (*User, errors.ErrorCode) {
	db := db.Connector()

	user := new(User)

	err := db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		return nil, errors.ErrorCode{Message: "Email not found", Code: 201}
	}

	if ok, _ := utils.CheckPasswordHash(password, user.Password); ok {
		return user, errors.ErrorCode{Message: "Login success", Code: 200}
	}

	return nil, errors.ErrorCode{Message: "Invalid credentials", Code: 500}
}
