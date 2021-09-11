package v1

import "fmt"

// GetInfo receives user information.
func GetInfo() (string, error) {
	var id string
	fmt.Println("Please enter your 'ID'")
	_, err := fmt.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
