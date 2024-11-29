package services

import (
	"bytes"
	"encoding/json"
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

// function to format JSON data
func FormatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
