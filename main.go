package main

import (
	"fmt"
	"os"
	"strconv"

	"go-cli-banking/services"
	"go-cli-banking/utils"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: create_account, transfer, add_deposit, accrue_interest")
		return
	}

	command := os.Args[1]

	err := utils.InitializeStorage()

	if err != nil {
		fmt.Println("Failed initializing storage:", err)
		return
	}

	switch command {

	case "create_account":

		if len(os.Args) != 4 {
			fmt.Println("Usage: create_account <name> <amount>")
			return
		}

		name := os.Args[2]

		amount, err := strconv.ParseFloat(os.Args[3], 64)

		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		services.CreateAccount(name, amount)

	case "transfer":

		if len(os.Args) != 5 {
			fmt.Println("Usage: transfer <from> <to> <amount>")
			return
		}

		from := os.Args[2]
		to := os.Args[3]

		amount, err := strconv.ParseFloat(os.Args[4], 64)

		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		services.Transfer(from, to, amount)

	case "add_deposit":

		if len(os.Args) != 4 {
			fmt.Println("Usage: add_deposit <name> <amount>")
			return
		}

		name := os.Args[2]

		amount, err := strconv.ParseFloat(os.Args[3], 64)

		if err != nil {
			fmt.Println("Invalid amount")
			return
		}

		services.AddDeposit(name, amount)

	case "accrue_interest":

		services.AccrueInterest()

	default:
		fmt.Println("Unknown command")
	}
}
