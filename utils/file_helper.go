package utils

import "os"

func InitializeStorage() error {

	folder := []string{
		"accounts",
		"deposits",
		"history",
	}

	for _, folder := range folder {

		err := os.MkdirAll(folder, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}