package main

import "fmt"

func main() {
	var Balance float64 = 1000

	fmt.Println("Welcome to the Go bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var Input int
	fmt.Print("Select an option: ")
	fmt.Scan(&Input)
	// wantsCheckBalance := Input == 1 // the variable will be boolean
	if Input == 1 {
		fmt.Println("Your balance is", Balance)
	} else if Input == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0")
			return
		}
		Balance += depositAmount
		fmt.Println("Balance updated! New Amount:", Balance)

	} else if Input == 3 {
		fmt.Print("Money to withdraw: ")
		var withdrawMoney float64
		fmt.Scan(&withdrawMoney)
		if withdrawMoney <= 0 || withdrawMoney >= Balance {
			fmt.Println("Invalid amount. Must be greater than 0 and less than your balance")
			return
		}

		Balance -= withdrawMoney
		fmt.Println("Your new Balance is: ", Balance)
	} else {
		fmt.Println("Goodbye")
	}
}
