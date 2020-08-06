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

var inputDateOfBirthLayout = "2006-01-02"
var genPrefPossibleState = []string{"male", "female","other"}

//Dater table schema
type Dater struct {
	Username                string `json:"username"`
	WingUsername            string `json:"wing_username"`
	SearchPerm             	string `json:"search_permission"`
	DateOfBirth             string `json:"date_of_birth"`
	ShortIntro            	string `json:"short_intro"`
	CurrentCity             string `json:"current_city"`
	JobRole               	string `json:"job_role"`
	EmploymentStatus        string `json:"employment_status"`
	StudyCollege            string `json:"study_college"`
	GenPref									string `json:"gender_preference"`
}

//GetDaterProfile gets the profile information for dater
func GetDaterProfile(c *gin.Context){
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

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
  var d = Dater{}

  if len(uname)>0 && len(uname)<=50{
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, wing_username, search_permission, date_of_birth, short_intro, current_city, job_role, employment_status, study_college, gender_preference from dater_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, uname).Scan(&d.Username, &d.WingUsername, &d.SearchPerm, &d.DateOfBirth, &d.ShortIntro, &d.CurrentCity, &d.JobRole, &d.EmploymentStatus, &d.StudyCollege, &d.GenPref)
		if err != nil {
      if err.Error() == "sql: no rows in result set"  {
        c.JSON(http.StatusOK, gin.H{"msg":d})
      	return
      }
			log.Fatal(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
		}
		d.DateOfBirth = d.DateOfBirth[:10]
		c.JSON(http.StatusOK, gin.H{"msg":d})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username length"})
	return
}

//PostDaterProfile will create and edit the wing_mate profile
func PostDaterProfile(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	var input Dater
	//var successMsg = map[string]string{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	empty:=Dater{}
	if input == empty{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty input"})
		return
	}
	
	genPref := p.StringInSlice(input.GenPref,genPrefPossibleState)
	if genPref != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gender_preference field"})
		return
	}

	_, err := time.Parse(inputDateOfBirthLayout, input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dob format"})
		return
	}
	
	log.Println(input.Username)
	if len(input.Username)>0 && len(input.Username)<=50 {

		sqlStatement := `
		INSERT INTO dater_profile (username, wing_username, search_permission, date_of_birth, short_intro, current_city, job_role, employment_status, study_college, gender_preference)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
				gender_preference = EXCLUDED.gender_preference,
				last_modified_on= now()`

		_, err := m.MyDB.Exec(sqlStatement, input.Username, input.WingUsername, input.SearchPerm, input.DateOfBirth, input.ShortIntro, input.CurrentCity, input.JobRole, input.EmploymentStatus, input.StudyCollege, input.GenPref)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg":"Success"})
		return

	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid length of username"})
	return
}
