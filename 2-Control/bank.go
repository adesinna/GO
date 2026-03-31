package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// The filename where we store our data
const accountBalanceFile = "balance.txt"

// writeBalanceToFile converts the float to a string and saves it to the disk
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		fmt.Println("Warning: Could not save balance to file.")
	}
}

// getBalanceFromFile reads the file, converts bytes to string, then string to float
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		// If the file doesn't exist, return 0 and a specific error message
		return 0, errors.New("no previous balance found. Starting fresh")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("error: balance file is corrupted")
	}

	return balance, nil
}

func addMoney(amount, balance float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("the amount entered is invalid")
	}
	return balance + amount, nil
}

func withdrawMoney(amount, balance float64) (float64, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	if amount <= 0 {
		return balance, errors.New("withdrawal amount must be positive")
	}
	return balance - amount, nil
}

func main() {
	fmt.Println("Welcome to the Shanana Bank (SB)")

	// --- 1. INITIAL LOAD FROM FILE ---
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("System Note:", err)
		// If it's a fresh start, create the file with 0.00
		writeBalanceToFile(0)
	} else {
		fmt.Printf("Data loaded successfully. Welcome back!\n")
		fmt.Printf("Your current balance is: $%.2f\n", balance)
	}

	var choice int

UserLoop:
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit funds")
		fmt.Println("3. Withdraw funds")
		fmt.Println("4. Exit")
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
				writeBalanceToFile(balance) // SAVE to disk
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
				writeBalanceToFile(balance) // SAVE to disk
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