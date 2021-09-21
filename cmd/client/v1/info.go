package v1

import "fmt"

// GetInfo receives user information.
func GetInfo() (string, error) {
	var id string

	fmt.Println("Please enter your 'ID'")

	if _, err := fmt.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
