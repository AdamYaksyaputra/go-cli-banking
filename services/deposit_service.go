package services

import (
	"fmt"
	"os"
	"time"
)

func AddDeposit(name string, amount float64) {

	if amount <= 0 {
		fmt.Println("Amount must be greater than zero")
		return
	}

	filePath := "deposits/" + name + ".txt"

	file, err := os.OpenFile(
		filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Failed to open deposit file")
		return
	}

	defer file.Close()

	depositData := fmt.Sprintf(
		"%.2f|%s\n",
		amount,
		time.Now().Format(time.RFC3339),
	)

	_, err = file.WriteString(depositData)

	if err != nil {
		fmt.Println("Failed to save deposit")
		return
	}

	fmt.Printf(
		"Deposit %.2f added for %s\n",
		amount,
		name,
	)
}