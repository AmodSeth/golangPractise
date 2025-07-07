package main

import (
	"fmt" 
	"os"
)

//project goals
/*
1. convert this into a function
2. also validate user inputs with error
- no negative numbers
- not 0
3. read and write to the files

*/

const tax_rate = 0.5


func writetothefile(price string){
	new_price := price 
	err := os.WriteFile("current_price.txt", []byte(new_price) , 0644)
	if err != nil {
		panic(err)
	}

}
func readfromthefile()  {
	price , err := os.ReadFile("current_price.txt")
	if err != nil { panic(err) }
	//need to conver this into string
	fmt.Println(string(price))
}

func getuserInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)

	return input
}
func validateMiddleware(value int) bool {
	if value < 0 || value == 0 {
		return false
	} else {
		return true
	}

	
}

func calculate_financials(revenue , expenses , tax_rate int ) (int,int) {
	//EBT=Revenue−Expenses(excludingtax)
	earningBeforeTax := revenue - expenses

	// Taxes=EBT×Tax Rate,
	Taxes := earningBeforeTax * tax_rate

	// Net Income=EBT−Taxes 
	earningAfterTax  := earningBeforeTax - Taxes

	return earningBeforeTax , earningAfterTax

}

func main() {
	var earningBeforeTax , earningAfterTax int

	revenue := getuserInput("Please enter the revenue")
	expenses := getuserInput("Please enter the expenses")




	earningBeforeTax , earningAfterTax  = calculate_financials(revenue,expenses , tax_rate)

	fmt.Println("Earning Before Tax:",earningBeforeTax )
	fmt.Println("Earning After Tax:",earningAfterTax )
	
}