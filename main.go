package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserID   string `json:"user_id"`
	City     string `json:"city"`
	Password string `json:"password" binding : "required"`
}

var Data map[string]User //data is of map type user is custom type data

func main() {

	Data = make(map[string]User)
	r := gin.Default()
	setupRoutes(r)
	r.Run(":8080")
}

/*
	we have to make the routers for
	GET		= to fetch
	POST 	= to create new insert/data
	PUT and	= to update
	DELETE	= to delete
	so the method is r.GET (`method of request`,function made for getting the request)
					and so on for all POST , PUT and DELETE
*/

func setupRoutes(r *gin.Engine) {
	r.GET("/user/:user_id", GetUserByUserID)
	r.GET("/user2", GetAllUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:user_id", UpdateUser)
	r.DELETE("/user/:user_id", DeleteUser)
}

//comment changes
func GetUserByUserID(c *gin.Context) { //gin.context contains both http request and http response
	UserID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id_is_missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Pswd := c.GetHeader("password")
	if Pswd == "" {
		res := gin.H{
			"error": "password is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}

	user := getUserByID(UserID, Pswd)
	res := gin.H{
		"user": user,
	}
	c.JSON(http.StatusOK, res)

}

// GetAllUser function
func GetAllUser(c *gin.Context) {

	res := gin.H{ //res:=  creating a response
		"User": Data,
	}
	c.JSON(http.StatusOK, res) //sending response
	// return
}

// CreateUser POST
func CreateUser(c *gin.Context) { //gin.context contains both http request and http response

	reqBody := User{}
	err := c.Bind(&reqBody) // ....? foes it merges the data
	if err != nil {
		res := gin.H{ //always remeber to create a response
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if reqBody.UserID == "" { // if any one create empty user id
		res := gin.H{
			"error": "User id must not be empty ",
		}
		c.JSON(http.StatusBadRequest, res) //sending response
		return
	}

	if reqBody.Password == "" {
		res := gin.H{
			"error": "password must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Data[reqBody.UserID] = reqBody
	res := gin.H{
		"Sucesss": true,
		"User":    reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}

// UPDATE User PUT
func UpdateUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{ /*gin.H is defined as type H map[string]struct{}*/
			"error": "user_id is missing",
		}
		c.JSON(http.StatusBadRequest, res) // we have to give response status ok but the user_id is wrong so it
		return
	}

	Pswd, ok := c.Params.Get("password")
	if ok == false {
		res := gin.H{
			"error": "password is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}

	/*
		use the function arguments method (the first method),
		validate the HTTP parameters with c.Validation.Required("addr").Ok?
	*/

	reqBody := User{}
	// err := c.Bind(&reqBody)

	// if err != nil {
	// 	res := gin.H{
	// 		"error": err,
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	if reqBody.UserID == "" {
		res := gin.H{
			"error": "user_id can not be updated",
		}
		c.JSON(http.StatusBadRequest, res)
		return

	}

	if reqBody.UserID != userID {
		res := gin.H{
			"error": "user_id can not be updated",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if reqBody.Password != Pswd {
		res := gin.H{
			"error": "password does not match",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Data[userID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
	}
	c.JSON(http.StatusOK, res)

}

func DeleteUser(c *gin.Context) { //r.DELETE("/user/:user_id", DeleteUser)

	userID, ok := c.Params.Get("user_id")

	if ok == false {
		res := gin.H{ /*gin.H is defined as type H map[string]struct{}*/
			"error": "user id is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return // we have to give response status ok but the user_id is wrong so it
	}

	Pswd := c.GetHeader("password")

	if cheakPswd(userID, Pswd) == false {
		res := gin.H{
			"success": false,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	} else {

		delete_user := deleteUserByID(userID)
		// delete_user := deleteUserByID(userID)

		res := gin.H{
			"message": delete_user,
		}
		c.JSON(http.StatusOK, res)
		return
		// }

	}
}

// reqBody := User{}

// if reqBody.Password == "" {
// 	res := gin.H{
// 		"error": "password must not be empty",
// 	}
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }

// if reqBody.Password != Pswd {
// 	res := gin.H{
// 		"error": "password is incorrect",
// 	}
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }
