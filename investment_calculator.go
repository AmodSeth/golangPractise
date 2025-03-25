package main

import "fmt"
import "math"

func main() {
	//initally storing the value as float : generally type casting
	//adding explicit instruction
	var investmentVariable float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentVariable * math.Pow((1+expectedReturnRate/100), years)
	
	fmt.Print(futureValue)

}

//decalaring a variable : var




