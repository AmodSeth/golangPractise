package main 

import "fmt"


type myAvengers struct {
	name string 
	power uint8
	age int16
	origin_place
}

type origin_place struct{
	birth_place string 
}


func (p myAvengers) power_calculator() uint8 {
	return p.power - 4;
}


func main() {
	//printing after initialisation
	var a1 myAvengers ;
	fmt.Printf("this is the name default %v \n" , a1.name);

	a1.name = "captain America"
	a1.power = 10
	a1.age = 45
	a1.birth_place = "queens"
	fmt.Printf("this is the name %v \n" , a1.name);
	fmt.Printf("this is the age %v \n" , a1.age);	
	fmt.Printf("this is the power %v \n" , a1.power);
	fmt.Printf("this is the origin %v \n" , a1.birth_place);
	
	fmt.Print("this is when avengers fight \n", a1.power_calculator(a1.power))



	
}