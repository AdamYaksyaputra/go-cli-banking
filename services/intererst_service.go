package services

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func AccrueInterest() {

	files, err := os.ReadDir("deposits")

	if err != nil {
		fmt.Println("Failed to read deposits directory")
		return
	}

	for _, file := range files {

		filePath := "deposits/" + file.Name()

		depositFile, err := os.Open(filePath)

		if err != nil {
			fmt.Println("Failed to open deposit file")
			continue
		}

		scanner := bufio.NewScanner(depositFile)

		var updatedDeposits []string

		for scanner.Scan() {

			line := scanner.Text()

			parts := strings.Split(line, "|")

			if len(parts) != 2 {
				continue
			}

			amount, err := strconv.ParseFloat(parts[0], 64)

			if err != nil {
				continue
			}

			timestamp, err := time.Parse(
				time.RFC3339,
				parts[1],
			)

			if err != nil {
				continue
			}

			elapsedMinutes := time.Since(timestamp).Minutes()

			newAmount := amount * math.Pow(1.01, elapsedMinutes)

			updatedLine := fmt.Sprintf(
				"%.2f|%s",
				newAmount,
				time.Now().Format(time.RFC3339),
			)

			updatedDeposits = append(
				updatedDeposits,
				updatedLine,
			)
		}

		depositFile.Close()

		output := strings.Join(updatedDeposits, "\n")

		err = os.WriteFile(
			filePath,
			[]byte(output),
			0644,
		)

		if err != nil {
			fmt.Println("Failed to update deposits")
			continue
		}

		fmt.Printf(
			"Interest accrued for %s\n",
			file.Name(),
		)
	}

	fmt.Println("Interest accrual completed")
}
