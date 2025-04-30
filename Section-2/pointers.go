package main


import "fmt"

//pointers: that store value address instead of value
//why pointers:?   avoid unnecessary copying of values, Directly mutate the value
//&: to get the address of a variable
// *: to get the value behind the pointer

//nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.

func getAdultYears (age *int)  int {
	*age += 10
	return *age
}



func main() {
	age:= 23
	fmt.Println("Age before change:", age)

	//address of the pointer
	var agePointer *int
	agePointer = &age
	fmt.Println("Address of age:", agePointer)

	//value of the pointer
	fmt.Println("Value of agePointer:", *agePointer)

	//getAdultYears(age)
	fmt.Println("Age in adult years",getAdultYears(&age))



	printvalues(appUser)


}