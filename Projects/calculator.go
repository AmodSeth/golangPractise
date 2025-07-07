package main

import "fmt"


//first_number, last_number

func main() {

for {
	var first_number int
	var last_number int

	var operator string

	fmt.Println("Enter the First Number")
	fmt.Scan(&first_number)
	fmt.Println("Enter the Second Number")
	fmt.Scan(&last_number)
	fmt.Println("Enter the Operator Use & in operator to Exit")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Addition of 2 numbers is", first_number + last_number)
	
	case "-":
		fmt.Println("Subtraction of 2 numbers is", first_number - last_number)
	

	case "/":
		fmt.Println("Division of 2 numbers is", first_number / last_number)


	case "*":
		fmt.Println("Multiplication of 2 numbers is", first_number * last_number)

	
	case "&":
		{break} 

	default:
		fmt.Println("Please make a valid operation")

	}
	println(first_number,last_number,operator)
	
}
	
}