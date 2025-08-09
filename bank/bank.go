package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644) // file permissions , the owner of the file can read and write and the others just can read
}
func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
	var Balance = getBalanceFromFile()

	for {

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
		switch Input {
		case 1:
			fmt.Println("Your balance is", Balance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				//return
				continue
			}
			Balance += depositAmount
			fmt.Println("Balance updated! New Amount:", Balance)
			writeBalanceToFile(Balance)

		case 3:
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 || withdrawMoney >= Balance {
				fmt.Println("Invalid amount. Must be greater than 0 and less than your balance")
				continue // salta a la siguiente iteracion del bucle
			}

			Balance -= withdrawMoney
			fmt.Println("Your new Balance is: ", Balance)
			writeBalanceToFile(Balance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for chosing our bank")
			return // to end the function

		}
	}

}
