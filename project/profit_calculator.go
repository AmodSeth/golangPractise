package main

// 3 variables revenue , expenses, tax_rate

import "fmt"

func main() {
	var revenue int
	var expenses int
	var tax_rate int


	fmt.Print("enter the revenue : ")
	fmt.Scan(&revenue)
	fmt.Print("enter the expenses : ")
	fmt.Scan(&expenses)
	fmt.Print("enter the tax_rate : ")
	fmt.Scan(&tax_rate)

	//EBT=Revenue−Expenses(excludingtax)
	earningBeforeTax := revenue - expenses
	fmt.Println("earningBeforeTax", earningBeforeTax)

	// Taxes=EBT×Tax Rate
	Taxes := earningBeforeTax * tax_rate

	// Net Income=EBT−Taxes 
	earningAfterTax  := earningBeforeTax - Taxes
	fmt.Println("earningAfterTax",earningAfterTax)

	

}


