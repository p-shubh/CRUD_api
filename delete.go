package main

import "fmt"

func deleteUserByID(userID string) string {
	// user := User{}
	message := "invalid user"
	if user, ok := Data[userID]; ok {

		delete(Data, userID)
		message = "User Deleted"
		fmt.Println(user)
		return message
	}
	// fmt.Println(user)
	return message
}
