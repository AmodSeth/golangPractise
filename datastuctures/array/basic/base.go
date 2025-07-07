package main

import (
	"fmt"
	"reflect"
)


func main() {
	//declaring an array
	//var intArr1 []int;
	

	//declaring with the fix length
	//var intArr2 [3]int = [3]int{1,2,3}

	//using a short hand for it as well

	// var ages = [3]int{20,25,30}

	//creating arrays of string

	var names = [4]string{"yoshi","mario","peach","bowser"}






	//ny omiting the length value it has now become a slice
	// var intArr2 []int = []int{1,2,3,4}


	//apending into an slice
	var intSlice []int = [ ]int{4,5,6}
	//now to add into it we can have a append function
	var lenSlice []int = []int{6,7,8}
	//adding the single value to the slice
	intSlice = append(intSlice, 5)
	//adding multiple values to the slice using the slice operator
	intSlice = append(intSlice, lenSlice...)
 




	fmt.Println(reflect.TypeOf(intSlice))
	fmt.Println(len(intSlice))

	fmt.Println(names, len(names))
}

//slices are just wraps around the arrays to give a more general, powerful interface to the sequences of data

