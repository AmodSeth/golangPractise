package main

import "fmt"


//marker item
type Marker struct{
	name string
	color string
	Quantity int
}

//price function
type Price struct{
	perUnit int
	perDozen int
}

//creating a struct 

type price_calculator struct {}

func (price_calculator) Total_price (m Marker, p Price) int{
	if m.Quantity > 12 {
		return m.Quantity * p.perDozen
	}else {
		return m.Quantity * p.perUnit
	}

}




func main() {
	fmt.Print("One function or struct should have one reason to change")
	
	marker:= Marker{name:"cello", color: "blue", Quantity: 143}
	price:= Price{perUnit: 12, perDozen: 10}

	calculator := price_calculator{}

	total := calculator.Total_price(marker,price)

	fmt.Println("total_price",total)
}