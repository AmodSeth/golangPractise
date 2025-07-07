package main

import "fmt"

// if a var is decalared outside the file it is available to use by the whole program
var RandomRate string = "79897"

// creating a utility functions
func printstatement(text string) {
	fmt.Print(text)

}

//creating a function with a return value

func printStatementWithReturnValue(text1, text2 string) string {
	//say it contcat 2 strings
	newstring := text1 + text2 + RandomRate
	return newstring

}

// using an new return way smarter way
// amounts have the type int and i want these amount to be returned sperately so i will
// declare them in function definition  only and not inside the function just assign the value inside it
// and just return it
func normalmathsfunction(amount1, amount2 int) (am1 float64, am2 float64) {
	am1 = float64(amount1) + 16.6
	am2 = float64(amount2) + 16.76543
	return 
}

func main() {
	fmt.Println("creating a utility functions")
	printStatementWithReturnValue("hello", "world")
	printstatement("hello \n")
}
