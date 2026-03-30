package main

import (
	"errors"
	"fmt"
)

func loop_test() int64 {
    // sum is inferred as an 'int'
    sum := 0

    for i := range 11 {
        if i == 9 {
            fmt.Println("Skipping 9...")
            continue 
        }

        // Note: i will never reach 11 because the loop condition is i < 11
        if i == 10 { 
            fmt.Println("Reaching the end at 10")
            sum += i // decide if you want to add 10 before breaking
            break
        }

        fmt.Printf("Current number: %v\n", i)
        sum += i
    }

    // You must convert 'int' to 'int64' to match the return type
    return int64(sum)
}


// addMoney takes the amount and current balance, returning the new total.
func addMoney(amount, balance float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("The amount is entered is invalid!")
	}
	return balance + amount, nil
}

// withdrawMoney returns the new balance and an error if funds are insufficient.
func withdrawMoney(amount, balance float64) (float64, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	return balance - amount, nil
}

func main() {
	balance := 0.0
	var choice int

	fmt.Println("Welcome to the Shanana Bank (SB)")

	// We use a for loop so the menu keeps appearing until the user exits.
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit funds")
		fmt.Println("3. Withdraw funds")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your current balance is: $%.2f\n", balance)

		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			newBalance, err := addMoney(depositAmount, balance)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fmt.Printf("You deposited: $%.2f. New balance: $%.2f\n", depositAmount, balance)
			}

		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)

			// Handling both return values: the new balance and the potential error
			newBalance, err := withdrawMoney(withdrawAmount, balance)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fmt.Printf("You withdrew: $%.2f. New balance: $%.2f\n", withdrawAmount, balance)
			}

		} else if choice == 4 {
			fmt.Println("Thank you for using Shanana Bank. Goodbye!")
			break // Exits the for loop

		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}

	test := loop_test()

	fmt.Println(test)
}
