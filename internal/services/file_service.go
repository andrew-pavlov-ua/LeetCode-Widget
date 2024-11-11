package services

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	r, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("ReadFile: error reading file with path '%s': %e", path, err)
	}

	return string(r)
}
