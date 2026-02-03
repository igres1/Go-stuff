package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time  //this is a time struct
}

type Admin struct{
	email string
	password string
	User User
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
} // with the parentesis just after the keyword func its called a RECEVIER argument , its for attaching the function to an struct , so the function now is part of that struct

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func NewAdmin(email , password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName : "ADMIN",
			lastName: "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("wrong parameters ")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(), //this function call creats a struct
	}, nil
}