package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time //this is a time struct
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
} // with the parentesis just after the keyword func its called a RECEVIER argument , its for attaching the function to an struct , so the function now is part of that struct

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {

	userFirstName := getUserData("Please enter your 1st name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(), //this function call creats a struct
	} //if u omit a value for example u dont put the userbirthdate it will be initialized to an empty string
	//appUser = user{}  //creates a empty struct

	appUser.outputUserDetails()
	appUser.clearUserName() //golang makes sure that a pointer of the struct is passed
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value

}
