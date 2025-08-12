//arrays
package main 
//Fixed length //same type (the default is set to default value of that data type )// indexable
import "fmt"


//slices are the wrapper around arrays


func main() {
	var myArray [3]int32; //0 is gonna print if no default is gonna be printed

	// fmt.Println(myArray)
	myArray[1]=1234;
	// fmt.Println(myArray);
	map_function();

}


func map_function(){
	// var myMap map[string]int = make(map[string]int);
	var myMap2 = map[string]uint8{"Adam":8 , "kritika":16};
	//if it does not return anything ok will be false and 
	var age , ok = myMap2["hello"];
	if ok {
		fmt.Println("age fount" , age)
	}else {
		fmt.Println("age not found");
	}
}