package main

import "fmt"


func main (){
	var mystring = "hello world"
	//to edit a string convert it into a rune
	var indexed = []rune(mystring);
	indexed[4] = 'g'

	mystring = string(indexed)
	fmt.Println("this is after rune", mystring)
	for s,a := range mystring{
		fmt.Printf("this is %v, %T \n", s,a);
	}
}
