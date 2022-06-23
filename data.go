package main

func getUserByID(userID string) User {
	user := User{}
	if user, ok := Data[userID] ; ok {
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