package controllers

type wingmate struct {
	Username                   string `json:"username"`
	DaterUsername             string `json:"dater_username"`
	RelationshipToDater      string `json:"relationship_to_dater"`
	DateOfBirth              string `json:"date_of_birth"`
	IntroWingLine            string `json:"intro_wing_line"`
	CurrentCity               string `json:"current_city"`
}

func GetWingmateProfile(c *gin.Context){
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  var wing_username string
  var wing_mate = wingmate{}
  if err := c.BindJSON(&wing_username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  
  if len(wing_username)>0{
		//see if username/email exists, if not then insert in to the table
		sqlStatement := `Select username, dater_username, relationship_to_dater, date_of_birth, intro_wing_line, current_city from dater_profile where username=$1`
		err := m.MyDB.QueryRow(sqlStatement, wing_username).Scan(&wing_mate.Username, &wing_mate.DaterUsername, &wing_mate.RelationshipToDater, &wing_mate.DateOfBirth, &wing_mate.IntroWingLine, &wing_mate.CurrentCity)
		if err != nil {
      if err.Error() == "sql: no rows in result set"  {
        c.JSON(http.StatusOK, gin.H{"msg":wing_mate})
      	return
      }
			log.Fatal(err)
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  		return
		}

	}
    
}