package main

func getUserByID(userID string, Password string) User {
	user := User{}
	if user, ok := Data[userID]; ok {
		if user, ok := Data[Password]; ok {
			return user
		}
		return user
	}
	return user
}

/*
		type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	UserID string `json:"User_id"`
	City   string `json:"city"`
}
*/
