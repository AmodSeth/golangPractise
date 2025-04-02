package main

import "fmt"

//project goals
/*
1. convert this into a function
2. also validate user inputs with error
- no negative numbers
- not 0
3. 

*/

func calculate_financials() {
	//EBT=Revenue−Expenses(excludingtax)
	earningBeforeTax := revenue - expenses
	fmt.Println("earningBeforeTax", earningBeforeTax)

	// Taxes=EBT×Tax Rate
	Taxes := earningBeforeTax * tax_rate

	// Net Income=EBT−Taxes 
	earningAfterTax  := earningBeforeTax - Taxes
	fmt.Println("earningAfterTax",earningAfterTax)
}

func getuserInput(){

}

func main() {

	
}