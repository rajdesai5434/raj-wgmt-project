package tests

import(
  "encoding/json"
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
  c "github.com/rajdesai5434/mah-cool-project/wmbe/controllers"
)

func TestGetWingmateProfileBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var expDict = c.Wingmate{}
  expDict.Username="rad5434"
	expDict.DaterUsername ="kevin"
  expDict.RelationshipToDater ="friend"
	expDict.DateOfBirth ="1997-01-05"
  expDict.IntroWingLine ="k..k..kevin"
  expDict.CurrentCity ="Sunnyvale"
  expResponse := gin.H{"msg":expDict}


  w := PerformGetRequestSingleParam(Router, "GET", "/api/wingmate_profile","username","rad5434")
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]c.Wingmate

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)
  var actDict = actResponse["msg"]

  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
  assert.Equal(t, expDict,actDict)

}

func TestGetWingmateProfileTrueCaseWrongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var expDict = c.Wingmate{}
  expResponse := gin.H{"msg": expDict}


  w := PerformGetRequestSingleParam(Router, "GET", "/api/wingmate_profile","username","poopy")
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]c.Wingmate

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"],actResponse["msg"])

}

func TestGetWingmateProfileFalseCaseNoUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  expResponse := gin.H{"error": "No Username"}

  w := PerformGetRequestSingleParam(Router, "GET", "/api/wingmate_profile","username","")
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"],actResponse["error"])

}

func TestGetWingmateProfileFalseCaseLongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  expResponse := gin.H{"error": "Invalid username length"}

  w := PerformGetRequestSingleParam(Router, "GET", "/api/wingmate_profile","username","dadsdsdsdfadsfasdfasdfasdfadfasdfasdfasdfaasdfasdfasdfdsdsdsfsdfsdfdfsdfds")
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"],actResponse["error"])

}

func TestPostWingmateProfileBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "kevin",
    "dater_username": "kal",
    "relationship_to_dater": "friend",
    "date_of_birth": "1997-01-04",
    "intro_wing_line": "Kalpi... kuthrapli from India",
    "current_city": "Matevale"
    }`)

  expResponse := gin.H{"msg":"Success"}


  w := PerformPostRequest(Router, "POST", "/api/wingmate_profile", inputStr)
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)


  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
}

func TestPostWingmateProfileFalseCaseWrongInput (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "kevin",
    "dater_username": "kal",
    "relationship_to_dater": "friend",
    "date_of_birth": "poop",
    "intro_wing_line": "Kalpi... kuthrapli from India",
    "current_city": "Matevale"
    }`)

  expResponse := gin.H{"error":"Invalid dob format"}


  w := PerformPostRequest(Router, "POST", "/api/wingmate_profile", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)


  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestPostWingmateProfileInvalidCaseNoBody (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{}`)

  expResponse := gin.H{"error":"Empty input"}


  w := PerformPostRequest(Router, "POST", "/api/wingmate_profile", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)


  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}