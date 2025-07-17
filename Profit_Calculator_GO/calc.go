package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxes float64

	revenue = readUserInput("revenue: ")

	expenses = readUserInput("Enter Expenses: ")

	taxes = readUserInput("Enter taxes to pay: ")

	EBT, profit, taxes := calculate(revenue, expenses, taxes)

	fmt.Printf("EBT = %3f\n", EBT)
	fmt.Printf("Profit = %3f\n", profit)
	fmt.Printf("Taxes= %3f", taxes)
}
func readUserInput(info string) float64 {
	var Input float64
	fmt.Print(info)
	fmt.Scan(&Input)
	return Input

}
func calculate(revenue, expenses, taxes float64) (float64, float64, float64) {
	EBT := revenue - expenses
	profit := EBT - (EBT * taxes / 100)
	ratio := EBT / profit
	return EBT, profit, ratio
}
