package main

import "fmt"

func main() {
	//using for loops
	//go only has for loops no while
	for i := 0; i<2 ;i++{
		fmt.Println("hello world")
	}
	//using flexible for loop for infinite loop
	//just use keyword for and wrap it in the loop 
	flag := true
	for  flag {
		fmt.Println("this is true")
	}
	
}