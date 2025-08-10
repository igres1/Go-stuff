package main

import (
	"errors"
	"fmt"
	"os"
)

const CalcFile = "calcFile.txt"

func main() {

	revenue, err := readUserInput("revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := readUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxes, err := readUserInput("Enter taxes to pay: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	EBT, profit, taxes := calculate(revenue, expenses, taxes)
	result := fmt.Sprintf("The EBT is %.2f , the profit is %.2f and the taxes are %.2f\n", EBT, profit, taxes)
	os.WriteFile(CalcFile, []byte(result), 0644)

	fmt.Printf("EBT = %.1f\n", EBT)
	fmt.Printf("Profit = %.1f\n", profit)
	fmt.Printf("Taxes= %.1f", taxes)
}
func readUserInput(info string) (float64, error) {
	var Input float64
	fmt.Print(info)
	fmt.Scan(&Input)
	if Input <= 0 {
		return 0, errors.New("The value can not be equal or less than 0")

	}
	return Input, nil

}
func calculate(revenue, expenses, taxes float64) (float64, float64, float64) {
	EBT := revenue - expenses
	profit := EBT - (EBT * taxes / 100)
	ratio := EBT / profit
	return EBT, profit, ratio
}
