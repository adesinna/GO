package main

import (
	"fmt"
)

func getAdultYear(age *int) int  {
	return *age - 18
}
func main()  {
	age := 32 
	// adult_years := getAdultYear(age) // the age will copy the 32 from the age and use it then dispose 

	userAge := &age

	adult_years := getAdultYear(userAge)

	fmt.Println("Age:", age)
	fmt.Println(adult_years)
}

