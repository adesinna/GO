package main

import (
	"fmt"
	"math"
)

func calculateCompoundInterest(investment, rate, years float64) float64 {
	inflationRate := 1.8
	compoundedAmount := investment * math.Pow(1+(rate/100), years)
	return compoundedAmount / math.Pow(1+(inflationRate/100), years)
}

func main() {
	var investmentAmount float64 = 10000
	var expectedReturnRate float64 = 5.5
	var years float64 = 10

	futureValue := calculateCompoundInterest(investmentAmount, expectedReturnRate, years)

	fmt.Println("This is the value of the first calculation")
	fmt.Printf("Future Value: $%.2f\n", futureValue)

	investmentAmount = 2370.36
	expectedReturnRate = 15.3
	years = 35

	futureValue = calculateCompoundInterest(investmentAmount, expectedReturnRate, years)

	fmt.Println("This is the value of the second calculation")
	fmt.Printf("Future Value: $%.2f\n", futureValue)

	// User input
	fmt.Println("--- Compound Interest Calculator (Inflation Adjusted) ---")

	fmt.Print("Enter Initial Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Annual Return (e.g., 5.5): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Number of Years: ")
	fmt.Scan(&years)

	futureValue = calculateCompoundInterest(investmentAmount, expectedReturnRate, years)

	fmt.Println("---------------------------------------------------------")
	fmt.Printf("Adjusted Future Value (in today's dollars): $%.2f\n", futureValue)

	name := "Gopher"
	age := 12

	// Sprintf doesn't print anything to the screen.
	// It creates a new string and saves it to the variable 'message'.
	message := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)

	// Now you can use that string whenever you want!
	fmt.Println(message)

}
