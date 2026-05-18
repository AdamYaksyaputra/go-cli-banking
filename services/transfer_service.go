package services

import "fmt"

func Transfer(from string, to string, amount float64) {

	if amount <= 0 {
		fmt.Println("Amount must be greater than zero")
		return
	}

	fmt.Printf(
		"Transfer %.2f from %s to %s\n",
		amount,
		from,
		to,
	)
}