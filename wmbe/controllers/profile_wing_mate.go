package controllers

import (
	"log"
	//"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
	p "github.com/rajdesai5434/mah-cool-project/wmbe/pkg"
)

var relToDaterPossibleState = []string{"friend", "relative","other"}

//Wingmate table schema
type Wingmate struct {
	Username                string `json:"username"`
	DaterUsername           string `json:"dater_username"`
	RelationshipToDater     string `json:"relationship_to_dater"`
	DateOfBirth             string `json:"date_of_birth"`
	IntroWingLine           string `json:"intro_wing_line"`
	CurrentCity             string `json:"current_city"`
}


//GetWingmateProfile gets the profile information for wing_mates
func GetWingmateProfile(c *gin.Context){
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  var Wingmate = Wingmate{}
	var uname string

	if len(c.Request.URL.Query())>0{
		var reqLen = len(c.Request.URL.Query()["username"])
		if reqLen>0 && (c.Request.URL.Query()["username"][0]!=""){
			uname=c.Request.URL.Query()["username"][0]
		} else{
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Username"})
			return
		}
	} else{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Username"})
		return
	}
	//uname:=c.Request.URL.Query()["username"][0]
  if len(uname)>0 && len(uname)<=50{
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, dater_username, relationship_to_dater, date_of_birth, intro_wing_line, current_city from wing_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, uname).Scan(&Wingmate.Username, &Wingmate.DaterUsername, &Wingmate.RelationshipToDater, &Wingmate.DateOfBirth, &Wingmate.IntroWingLine, &Wingmate.CurrentCity)
		if err != nil {
      if err.Error() == "sql: no rows in result set"  {
        c.JSON(http.StatusOK, gin.H{"msg":Wingmate})
      	return
      }
			log.Fatal(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
		}
		Wingmate.DateOfBirth = Wingmate.DateOfBirth[:10]
		c.JSON(http.StatusOK, gin.H{"msg":Wingmate})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username length"})
	return
}

//PostWingmateProfile will create and edit the wing_mate profile
func PostWingmateProfile(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	var input Wingmate
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empty:=Wingmate{}
	if input == empty{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty input"})
		return
	}

	_, err := time.Parse(inputDateOfBirthLayout, input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dob format"})
		return
	}

	relToDater := p.StringInSlice(input.RelationshipToDater,relToDaterPossibleState)
	if relToDater != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid relationship_to_dater field"})
		return
	}

	if len(input.Username)>0 && len(input.Username)<=50{

		sqlStatement := `
		INSERT INTO wing_profile (username, dater_username, relationship_to_dater, date_of_birth, intro_wing_line, current_city)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (username)
		DO
			UPDATE SET
				dater_username = EXCLUDED.dater_username,
				relationship_to_dater = EXCLUDED.relationship_to_dater,
				date_of_birth = EXCLUDED.date_of_birth,
				intro_wing_line = EXCLUDED.intro_wing_line,
				current_city = EXCLUDED.current_city,
				last_modified_on= now()`

		_, err := m.MyDB.Exec(sqlStatement, input.Username, input.DaterUsername, input.RelationshipToDater, input.DateOfBirth, input.IntroWingLine, input.CurrentCity)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg":"Success"})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username length"})
	return
}
