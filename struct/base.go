package main

import (
	"fmt"
	"time"

)

//creating a struct
//take input of first name , last name and birth date

type User struct {
	firstname string
	lastname string
	birthdate string
	createdAT time.Time
}


func NewUser(firstname, lastname, birthdate string) User {
	return User{
		firstname: firstname,
		lastname: lastname,
		birthdate: birthdate,
		createdAT: time.Now(),
	}
}


func (u  User) user_output () {
	fmt.Println("Your First Name is", u.firstname)
	fmt.Println("Your Birth Date is", u.birthdate)
	fmt.Println("Your Last Name is", u.lastname)
	fmt.Println("crated at", u.createdAT.Format("2006-01-02 15:04:05"))
}


//creating a clear user name function
func (u *User) clearUserName ()  {
	u.firstname = ""
	u.lastname = ""
	fmt.Println("user name was cleared");
}



func user_input (text string) string{
	fmt.Println(text)
	var input string;
	fmt.Scan(&input)

	return input
}



func main() {

	user_firstname := user_input("Enter your First Name")
	user_lastname := user_input("Enter your Last Name")
	user_birthdate := user_input("Enter your Birth Date in YYYY-MM-DD format")

	user := NewUser(user_firstname, user_lastname, user_birthdate)
	
	user.user_output()
	//user.clearUserName()
	user.user_output()

}