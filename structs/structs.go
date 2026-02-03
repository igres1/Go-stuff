package main

import (
	"fmt"
	"learning-Go/structs/user"
)

func main() {

	userFirstName := getUserData("Please enter your 1st name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("example@example.com ", "yesyes")
	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	/*appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(), //this function call creats a struct
	} //if u omit a value for example u dont put the userbirthdate it will be initialized to an empty string
	//appUser = user{}  //creates a empty struct
	*/
	appUser.OutputUserDetails()
	appUser.ClearUserName() //golang makes sure that a pointer of the struct is passed
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value

}
