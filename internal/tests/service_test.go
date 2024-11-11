package testing

import (
	"testing"

	"cmd/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	filepath := "file_to_scan.txt"
	expectedFileString := "hello world!"
	actualFileString := services.ReadFile(filepath)

	assert.Equal(t, expectedFileString, actualFileString, "Expected parsed string equals 'hello world!'")
}
