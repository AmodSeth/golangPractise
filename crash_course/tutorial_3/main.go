package main

import "fmt"

func main() {
	var	dividend int ; 
	var remainder int ;

	fmt.Printf("enter the numerator");
	fmt.Scan(&dividend)

	fmt.Printf("enter the denominator");
	fmt.Scan(&remainder)

	result , res := divisor_function(dividend,remainder);
	fmt.Printf("result is %d, remainder is %d  \n", result , res);
	
}



func divisor_function(dividend,remainder int) (int, int){
	result := dividend / remainder
	rem := dividend % remainder
	return result , rem;
}