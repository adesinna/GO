package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// writeBalanceToFile converts the float to a string and saves it to the disk
func WriteBalanceToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%.2f", value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		fmt.Println("Warning: Could not save balance to file.")
	}
}

// getBalanceFromFile reads the file, converts bytes to string, then string to float
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		// If the file doesn't exist, return 0 and a specific error message
		return 0, errors.New("no previous value found. Starting fresh")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("error: file is corrupted")
	}

	return value, nil
}

