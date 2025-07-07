package main


import "fmt"

func helperfunction(text1, text2 string) (contcat string)  {
	contcat = text1 + text2
	return
}

func main (){
	fmt.Print(helperfunction("origin","objects"))
	return;
}3