package controllers

import (
	"log"
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)

type wingmate struct {
	Username                   string `json:"username"`
	DaterUsername             string `json:"dater_username"`
	RelationshipToDater      string `json:"relationship_to_dater"`
	DateOfBirth              string `json:"date_of_birth"`
	IntroWingLine            string `json:"intro_wing_line"`
	CurrentCity               string `json:"current_city"`
}

type dater struct {
	Username                   string `json:"username"`
	WingUsername             string `json:"wing_username"`
	SearchPerm             string `json:"search_permission"`
	DateOfBirth              string `json:"date_of_birth"`
	ShortIntro            string `json:"short_intro"`
	CurrentCity               string `json:"current_city"`
	JobRole               string `json:"job_role"`
	EmploymentStatus               string `json:"employment_status"`
	StudyCollege               string `json:"study_college"`
}


//GetWingmateProfile gets the profile information for wing_mates
func GetWingmateProfile(c *gin.Context){
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  var wingMate = wingmate{}

  if len(c.Param("username"))>0{
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, dater_username, relationship_to_dater, date_of_birth, intro_wing_line, current_city from wing_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, c.Param("username")).Scan(&wingMate.Username, &wingMate.DaterUsername, &wingMate.RelationshipToDater, &wingMate.DateOfBirth, &wingMate.IntroWingLine, &wingMate.CurrentCity)
		if err != nil {
      if err.Error() == "sql: no rows in result set"  {
        c.JSON(http.StatusOK, gin.H{"msg":wingMate})
      	return
      }
			log.Fatal(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
		}
		c.JSON(http.StatusOK, gin.H{"msg":wingMate})
		return
	}
}

//PostWingmateProfile will create and edit the wing_mate profile
func PostWingmateProfile(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	var input wingmate
	//var successMsg = map[string]string{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Username)>0 && len(input.DateOfBirth)>0 {

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
		//successMsg["username"]=input.Username
		//successMsg["appUseStatus"]=input.AppUseStatus
		c.JSON(http.StatusOK, gin.H{"msg":"Success"})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Fields"})
	return
}

//GetDaterProfile gets the profile information for dater
func GetDaterProfile(c *gin.Context){
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  var d = dater{}

  if len(c.Param("username"))>0{
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, wing_username, search_permission, date_of_birth, short_intro, current_city, job_role, employment_status, study_college from dater_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, c.Param("username")).Scan(&d.Username, &d.WingUsername, &d.SearchPerm, &d.DateOfBirth, &d.ShortIntro, &d.CurrentCity, &d.JobRole, &d.EmploymentStatus, &d.StudyCollege)
		if err != nil {
      if err.Error() == "sql: no rows in result set"  {
        c.JSON(http.StatusOK, gin.H{"msg":d})
      	return
      }
			log.Fatal(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
		}
		c.JSON(http.StatusOK, gin.H{"msg":d})
		return
	}
}

//PostDaterProfile will create and edit the wing_mate profile
func PostDaterProfile(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	var input dater
	//var successMsg = map[string]string{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Username)>0 && len(input.DateOfBirth)>0 {

		sqlStatement := `
		INSERT INTO dater_profile (username, wing_username, search_permission, date_of_birth, short_intro, current_city, job_role, employment_status, study_college)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (username)
		DO
			UPDATE SET
				wing_username = EXCLUDED.wing_username,
				search_permission = EXCLUDED.search_permission,
				date_of_birth = EXCLUDED.date_of_birth,
				short_intro = EXCLUDED.short_intro,
				current_city = EXCLUDED.current_city,
				job_role = EXCLUDED.job_role,
				employment_status = EXCLUDED.employment_status,
				study_college = EXCLUDED.study_college,
				last_modified_on= now()`

		_, err := m.MyDB.Exec(sqlStatement, input.Username, input.WingUsername, input.SearchPerm, input.DateOfBirth, input.ShortIntro, input.CurrentCity, input.JobRole, input.EmploymentStatus, input.StudyCollege)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//successMsg["username"]=input.Username
		//successMsg["appUseStatus"]=input.AppUseStatus
		c.JSON(http.StatusOK, gin.H{"msg":"Success"})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Fields"})
	return
}
