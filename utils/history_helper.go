package utils

import (
	"fmt"
	"os"
	"time"
)

func AddHistory(name string, message string) error {

	filePath := "history/" + name + ".txt"

	file, err := os.OpenFile(
		filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	log := fmt.Sprintf(
		"[%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		message,
	)

	_, err = file.WriteString(log)

	if err != nil {
		return err
	}

	return nil
}