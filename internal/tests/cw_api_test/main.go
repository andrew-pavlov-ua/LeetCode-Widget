package main

import (
	codewars_api "cmd/internal/cw_api"
	"fmt"
)

func main() {
	id := "Zorrix"
	data, err := codewars_api.GetUserProfile(id)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Data: ", data)
}
