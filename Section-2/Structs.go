package main

import (
	"fmt"
	"time"
)

//we use capital casing to make this variable avialable for the different files


//creating a new st 
type User struct{
	firstName string
	lastName string
	age int
	createdAt time.Time  
}

func printvalues( u User){
	fmt.Println(u.createdAt)
	fmt.Println(u.lastName)
	fmt.Println(u.firstName)
	fmt.Println(u.age)
}

func main() {
	userFirstName := "amod"
	userLastName := "seth"
	age := 23

	//creating an instance  of the struct using struct literal method 
	
	
	
	appUser := User{
		firstName: userFirstName ,
		lastName: userLastName,
		age : age,
		createdAt: time.Now(),
		
		}
		
	// note if order is the same we can omit the first key but it must remain the same order for sure

	appUser2 := User{
		userFirstName, 
		userLastName, 
		age, 
		time.Now(),

	}


	printvalues(appUser)
	printvalues(appUser2)


	return
}