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

func TestGetDaterProfileBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var expDict = c.Dater{}
  expDict.Username= "kal"
	expDict.WingUsername = "rad5434"
  expDict.SearchPerm = "true"
  expDict.DateOfBirth = "1995-04-06"
  expDict.ShortIntro = "I am Kalpi, people love me, women want to be with me, men want to be me"
  expDict.CurrentCity = "Poopvile"
  expDict.JobRole = "DE"
  expDict.EmploymentStatus = "Employed"
  expDict.StudyCollege = "PSU"
  expDict.GenPref = "other"
  expResponse := gin.H{"msg":expDict}


  w := PerformGetRequestSingleParam(Router, "GET", "/api/dater_profile","username","kal")
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]c.Dater

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)
  var actDict = actResponse["msg"]

  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
  assert.Equal(t, expDict,actDict)

}

func TestGetDaterProfileTrueCaseWrongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var expDict = c.Dater{}
  expResponse := gin.H{"msg": expDict}


  w := PerformGetRequestSingleParam(Router, "GET", "/api/dater_profile","username","poopy")
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]c.Dater

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"],actResponse["msg"])

}

func TestGetDaterProfileFalseCaseNoUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  expResponse := gin.H{"error": "No Username"}

  w := PerformGetRequestSingleParam(Router, "GET", "/api/dater_profile","username","")
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"],actResponse["error"])

}

func TestGetDaterProfileFalseCaseLongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  expResponse := gin.H{"error": "Invalid username length"}

  w := PerformGetRequestSingleParam(Router, "GET", "/api/dater_profile","username","dadsdsdsdfadsfasdfasdfasdfadfasdfasdfasdfaasdfasdfasdfdsdsdsfsdfsdfdfsdfds")
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"],actResponse["error"])

}

func TestPostDaterProfileBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "kal",
    "wing_username": "rad5434",
    "search_permission": "true",
    "date_of_birth": "1995-04-06",
    "short_intro": "I am Kalpi, people love me, women want to be with me, men want to be me",
    "job_role": "DE",
    "employment_status": "Employed",
    "current_city": "Poopvile",
    "study_college":"PSU",
    "gender_preference": "other"
    }`)

  expResponse := gin.H{"msg":"Success"}

  w := PerformPostRequest(Router, "POST", "/api/dater_profile", inputStr)
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
}

func TestPostDaterProfileFalseCaseWrongInput (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "kal",
    "wing_username": "rad5434",
    "search_permission": "true",
    "date_of_birth": "1995-04-",
    "short_intro": "I am Kalpi, people love me, women want to be with me, men want to be me",
    "job_role": "DE",
    "employment_status": "Employed",
    "current_city": "Poopvile",
    "study_college":"PSU",
    "gender_preference": "other"
    }`)

  expResponse := gin.H{"error":"Invalid dob format"}


  w := PerformPostRequest(Router, "POST", "/api/dater_profile", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)


  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestPostDaterProfileInvalidCaseNoBody (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{}`)

  expResponse := gin.H{"error":"Empty input"}


  w := PerformPostRequest(Router, "POST", "/api/dater_profile", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]string

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)


  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}
