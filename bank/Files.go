package main

import "os"
import "fmt"

func writebalancetofile() {
	current_balance := "2000"
	os.WriteFile("balance.txt", []byte(current_balance), 0644)
}

func main() {
	fmt.Println("method of printing this file")
	writebalancetofile()
	return

}