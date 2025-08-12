package main

import (
	"errors"
	"fmt"
	"strconv"
)


type user struct {
	my_username string
	my_age int
}
//attaching a function to the struct

func (u user) printuserdata(){
	fmt.Printf("Method Called : your username is %s \n", u.my_username)
	fmt.Printf("Method Called: your age is %d \n", u.my_age)
}



//creatiing a user utility
func newUser (username string , age int) (*user, error) {
	if (username == "" || age == 0){
		return nil, errors.New("username and age can't be empty");

	}
	return &user{
		my_username: username,
		my_age: age,
	},nil
}

func getuserdata(text string) string {
	fmt.Println(text);
	var input string;
	fmt.Scanln(&input)

	return input
}


//constructor function




// func outputuserdetails(text1 string , text2 int) {
// 	fmt.Printf("your username is %s \n", text1)
// 	fmt.Printf("your age is %d \n", text2)

// }


func outputuserdetails(u *user)  {
	fmt.Printf("your username is %s \n", u.my_username)
	fmt.Printf("your age is %d \n", u.my_age)

}


func main() {
	username := getuserdata("Please enter your name?");
	ageStr := getuserdata("Please enter your age?");

	age,err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("please enter the valid age input")
		return

	}

	var appUser user
 

	//creating a struct literal
	appUser, error  := newUser(username,age)
	if err != nil {
		fmt.Println("Error" , error)
		return
	}
	



	//send the addusers to the appUser
	outputuserdetails(&appUser);
	//calling the method to the user
	appUser.printuserdata()

}


