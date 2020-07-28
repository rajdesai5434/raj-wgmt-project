package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)

type createNewUser struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	AppUseStatus string `json:"appUseStatus"`
}

//1. Check if the given email or username exist, if it does return asking for newer values
//2. After creating the new user, return the username in the api return
func CreateNewUserPost(c *gin.Context) {
	var input createNewUser
	var row = createNewUser{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username != "" && input.Password != "" && input.Email != "" && input.Fname != "" && input.Lname != "" && input.AppUseStatus != "" {
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, email from user_profile where username=$1 or email=$2`
		err := m.MyDB.QueryRow(sqlStatement, input.Username, input.Email).Scan(&row.Username, &row.Email)
		if err != nil && err.Error() != "sql: no rows in result set" {
			log.Fatal(err)
		}

		if row.Username == "" && row.Email == "" { //create a new entry
			sqlStatement = `
		   INSERT INTO user_profile (username, password, first_name, last_name, email, app_use_status, created_on)
		   VALUES ($1, $2, $3, $4, $5, $6, $7)`

			_, err = m.MyDB.Exec(sqlStatement, input.Username, input.Password, input.Fname, input.Lname, input.Email, input.AppUseStatus, time.Now())

			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, input.Username)
		} else {
			c.JSON(http.StatusConflict, "User or Email already in use")
		}

	}
}
