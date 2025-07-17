package main

import "fmt"

func main() {

	var revenue float32
	var expenses float32
	var taxes float32

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter  taxes to pay in decimal: ")
	fmt.Scan(&taxes)

	EBT := revenue - expenses
	profit := EBT - (EBT * taxes / 100)
	ratio := EBT / profit

	fmt.Print(ratio)

}
