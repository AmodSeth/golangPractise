package main

import "fmt"

func main() {
	fmt.Println("welcome to the bank")
	fmt.Println("what do you want to choose")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Scan(&choice)
	// not paranthesis
	if choice == 4 {
		fmt.Println("program has been exited")
		return 
	} else if choice >= 1 && choice <= 3 {
		//proceed
		fmt.Println("we can proceed")
	
	}else {
		fmt.Println("Invalid choice Please try again")
	}
}