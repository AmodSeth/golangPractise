// package main

// import "fmt"
// import "math"

// func main() {
// 	//initally storing the value as float : generally type casting
// 	//adding explicit instruction
// 	var investmentVariable float64 = 1000
// 	var expectedReturnRate = 5.5
// 	var years float64 = 10

// 	var futureValue = investmentVariable * math.Pow((1+expectedReturnRate/100), years)

// 	fmt.Print(futureValue)

// }

// //decalaring a variable : var

package main

import "fmt"

func main() {
	const inflation_rate = 7.7
	var mysalary int = 50000


	var incrementAmount int 

	//this is how to take the scanned input
	fmt.Scan(&incrementAmount)

	fmt.Println("incrementAmount", incrementAmount)

	finalAmountBeforeInflation := mysalary + incrementAmount

	fmt.Println("finalAmountBeforeInflation", finalAmountBeforeInflation)




}


