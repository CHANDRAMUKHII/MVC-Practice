package view

import (
	"fmt"
)

func GetUserChoice() string {
	fmt.Println("Welcome, admin!")
	var key string
	fmt.Println("What would you like to do? \n 1.Create new patient \n 2.Update patient details \n 3.Search for patient \n 4.Delete patient details \n 5.Exit")
	fmt.Scan(&key)

	return key
}
