package controllers

import (
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"database/sql"
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

type authenticateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppUseStatus string `json:"appUseStatus"`
}

//CreateNewUserPost creates a new entry in db for a given username and email.
func CreateNewUserPost(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	var input createNewUser
	var successMsg = map[string]string{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Username)>0 && len(input.Password)>0  && len(input.Email)>0 && len(input.Fname)>0 && len(input.Lname)>0 {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
		if err!=nil{
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sqlStatement := `
		   INSERT INTO user_profile (username, password, first_name, last_name, email, app_use_status)
		   VALUES ($1, $2, $3, $4, $5, $6)`

			_, err = m.MyDB.Exec(sqlStatement, input.Username, string(hashedPassword), input.Fname, input.Lname, input.Email, input.AppUseStatus)
			if err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			successMsg["username"]=input.Username
			successMsg["appUseStatus"]=input.AppUseStatus
			c.JSON(http.StatusOK, gin.H{"msg":successMsg})
			return

	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Fields"})
	return
}

//ApproveUserSignIn approves if the user exists and entered the right password
func ApproveUserSignIn (c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var input authenticateUser
	var response = authenticateUser{}
	var successMsg = map[string]string{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Username)>0 && len(input.Password)>0{

		sqlStatement := `Select password, app_use_status from user_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, input.Username).Scan(&response.Password,&response.AppUseStatus)
		if err != nil{
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(input.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	successMsg["username"]=input.Username
	successMsg["appUseStatus"]=response.AppUseStatus
	c.JSON(http.StatusOK, gin.H{"msg":successMsg})
	return

	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty Username or Password"})
	return
}
