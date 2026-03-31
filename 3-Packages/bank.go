package main

import (
	"fmt"
	"packages/fileops"
	"github.com/Pallinder/go-randomdata"

)


func main() {
	const accountBalanceFile = "balance.txt"

	fmt.Println("Welcome to the Shanana Bank (SB)")
	fmt.Println("Reach us 24/7 @",randomdata.PhoneNumber())

	// --- 1. INITIAL LOAD FROM FILE ---
	balance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("System Note:", err)
		// If it's a fresh start, create the file with 0.00
		fileops.WriteBalanceToFile(0, accountBalanceFile)
	} else {
		fmt.Printf("Data loaded successfully. Welcome back!\n")
		fmt.Printf("Your current balance is: $%.2f\n", balance)
	}

	var choice int

UserLoop:
	for {
		presentOptions() // thiis from welcome_menu.go which is part of this main package
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Current balance: $%.2f\n", balance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)

			newBalance, err := addMoney(depositAmount, balance)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fileops.WriteBalanceToFile(balance, accountBalanceFile) // SAVE to disk
				fmt.Printf("Successfully deposited $%.2f. New balance: $%.2f\n", depositAmount, balance)
			}

		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawAmount)

			newBalance, err := withdrawMoney(withdrawAmount, balance)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fileops.WriteBalanceToFile(balance, accountBalanceFile) // SAVE to disk
				fmt.Printf("Successfully withdrew $%.2f. New balance: $%.2f\n", withdrawAmount, balance)
			}

		case 4:
			fmt.Println("Thank you for choosing SB. Goodbye!")
			break UserLoop

		default:
			fmt.Println("Invalid option. Please choose 1-4.")
		}
	}
}