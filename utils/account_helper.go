package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadBalance(name string) (float64, error) {

	filePath := "accounts/" + name + ".txt"

	data, err := os.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	content := strings.TrimSpace(string(data))

	balance, err := strconv.ParseFloat(content, 64)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func WriteBalance(name string, amount float64) error {

	filePath := "accounts/" + name + ".txt"

	data := fmt.Sprintf("%.2f", amount)

	err := os.WriteFile(
		filePath,
		[]byte(data),
		0644,
	)

	if err != nil {
		return err
	}

	return nil
}