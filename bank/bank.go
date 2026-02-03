package main

import (
	"bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var Balance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("Can't continue sorry")
	}
	fmt.Println("Welcome to the Go bank")
	fmt.Println("reach us 24/7", randomdata.PhoneNumber())
	for {

		Greetings()

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
			fileops.WriteFloatToFile(Balance, accountBalanceFile)

		case 3:
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 || withdrawMoney >= Balance {
				fmt.Println("Invalid amount. Must be greater than 0 and less than your balance")
				continue // salta a la siguiente iteracion del bucle
			}

			Balance -= withdrawMoney
			fmt.Println("Your new Balance is: ", Balance)
			fileops.WriteFloatToFile(Balance, accountBalanceFile)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for chosing our bank")
			return // to end the function

		}
	}

}
