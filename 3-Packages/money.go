package main

import "errors"


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
