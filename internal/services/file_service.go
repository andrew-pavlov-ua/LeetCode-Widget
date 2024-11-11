package services

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	r, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	return string(r)
}
