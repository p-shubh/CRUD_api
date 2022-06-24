package main

func cheakPswd(UserID string, Pswd string) bool {

	var result bool = false
	if Pswd == Data[UserID].Password {
		result = true
	}

	// }

	return result

}
