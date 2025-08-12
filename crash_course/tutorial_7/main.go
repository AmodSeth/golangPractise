package main

import "fmt"

func main() {
	myArray := [5]int{1,2,3,4,5};
	fmt.Printf("memory location of the orignal array %p \n", &myArray);
	new_array := square(&myArray);
	

	for i := range new_array{
		fmt.Println(new_array[i]);
	}
}


func square(*myArray [5]int) [5]int {
	fmt.Printf("memory location of the orignal array %p \n", &arr);
	for i := range arr {
		arr[i] = arr[i] * arr[i];
	}
	return arr;
}