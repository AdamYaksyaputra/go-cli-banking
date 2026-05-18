package services

import (
	"fmt"

	"go-cli-banking/utils"
)

func Transfer(from string, to string, amount float64) {

	if amount <= 0 {
		fmt.Println("Amount must be greater than zero")
		return
	}

	if from == to {
		fmt.Println("Cannot transfer to same account")
		return
	}

	fromBalance, err := utils.ReadBalance(from)

	if err != nil {
		fmt.Println("Source account not found")
		return
	}

	toBalance, err := utils.ReadBalance(to)

	if err != nil {
		fmt.Println("Destination account not found")
		return
	}

	if fromBalance < amount {
		fmt.Println("Insufficient balance")
		return
	}

	newFromBalance := fromBalance - amount
	newToBalance := toBalance + amount

	err = utils.WriteBalance(from, newFromBalance)

	if err != nil {
		fmt.Println("Failed to update source balance")
		return
	}

	err = utils.WriteBalance(to, newToBalance)

	if err != nil {
		fmt.Println("Failed to update destination balance")
		return
	}

	fmt.Printf(
		"Transfer %.2f from %s to %s successful\n",
		amount,
		from,
		to,
	)

	utils.AddHistory(
		from,
		fmt.Sprintf(
			"Transfer sent to %s: %.2f",
			to,
			amount,
		),
	)

	utils.AddHistory(
		to,
		fmt.Sprintf(
			"Transfer received from %s: %.2f",
			from,
			amount,
		),
	)
}
