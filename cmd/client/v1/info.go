package v1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetInfo receives user information.
func GetInfo() (string, error) {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	_, err := writer.WriteString("Please enter your 'ID'\n")
	if err != nil {
		return "", fmt.Errorf("write to CL: %w", err)
	}

	if err = writer.Flush(); err != nil {
		return "", fmt.Errorf("write to CL: %w", err)
	}

	id, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("read from CL: %w", err)
	}

	id = strings.ReplaceAll(id, "\n", "")

	return id, nil
}
