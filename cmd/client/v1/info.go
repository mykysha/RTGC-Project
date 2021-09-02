package v1

import "fmt"

// Gets userID.
func GetInfo() string {
	var id string
	fmt.Println("Please enter your 'ID'")
	fmt.Scan(&id)

	return id
}
