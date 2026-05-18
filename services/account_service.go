package services

import (
	"fmt"
	"os"
)

func CreateAccount(name string, amount float64) {

	if amount <= 0 {
		fmt.Println("Amount must be greater than zero")
		return
	}

	filePath := "accounts/" + name + ".txt"

	_, err := os.Stat(filePath)

	if err == nil {
		fmt.Println("Account already exists")
		return
	}

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Failed to create account")
		return
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%.2f", amount))

	if err != nil {
		fmt.Println("Failed to write balance")
		return
	}

	fmt.Printf("Account %s created successfully with balance %.2f\n", name, amount)
}