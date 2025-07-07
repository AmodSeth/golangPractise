package main

import (
	"fmt"
	"os"
	"strconv"
)

func writebalancetofile() {
	current_balance := "2000"
	os.WriteFile("balance.txt", []byte(current_balance), 0644)
}

func readbalancefromfile(){
	data, _ := os.ReadFile(balance.txt)
	balancetext := string(data)
	balance ,  _ :=   strconv.ParseFloat(balancetext, 64)
	return balance

}

func main() {
	fmt.Println("method of printing this file")
	writebalancetofile()
	return

}