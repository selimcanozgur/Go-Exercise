package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter you first name: ")
	userLastName := getUserData("Please enter you last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser,err := user.New(userFirstName,userLastName,userBirthDate)

	
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("selimozgur26@gmail.com","pass1234")

	admin.GetUserDetail()
	admin.ClearUser()
	admin.GetUserDetail()

	appUser.GetUserDetail()
	appUser.ClearUser()
	appUser.GetUserDetail()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}