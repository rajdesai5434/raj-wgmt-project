package tests

import(
  "encoding/json"
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)

func TestApproveUserSignInBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"rad5434","password":"GO"}`)

  var expDict = make(map[string]interface{})
  expDict["username"]="rad5434"
	expDict["appUseStatus"]="wing_mate"
  expResponse := gin.H{"msg":expDict}


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]interface{}
  var actDict map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)
  actDict = actResponse["msg"].(map[string]interface{})

  val1, exists1 := actDict["username"]
  val2, exists2 := actDict["appUseStatus"]

  assert.Nil(t, err)
  assert.True(t, exists1)
  assert.True(t, exists2)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
  assert.Equal(t, expDict["username"], val1)
  assert.Equal(t, expDict["appUseStatus"], val2)
}

func TestApproveUserSignInFalseCasePwd (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"rad5434","password":"POOP"}`)

  expResponse := gin.H{"error":"crypto/bcrypt: hashedPassword is not the hash of the given password"}
  var actResponse map[string]string


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusUnauthorized, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestApproveUserSignInFalseCaseUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"rad","password":"GO"}`)

  expResponse := gin.H{"error":"sql: no rows in result set"}
  var actResponse map[string]string

  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusUnauthorized, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestApproveUserSignInInvalidCaseUnixPwd (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"rad5434","password":"भारत"}`)

  expResponse := gin.H{"error":"crypto/bcrypt: hashedPassword is not the hash of the given password"}
  var actResponse map[string]string


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusUnauthorized, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestApproveUserSignInInvalidCaseUnixUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"भारत","password":"GO"}`)

  expResponse := gin.H{"error":"sql: no rows in result set"}
  var actResponse map[string]string


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusUnauthorized, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestApproveUserSignInInvalidCaseEmptyUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"password":"GO"}`)

  expResponse := gin.H{"error":"Invalid username/password length"}
  var actResponse map[string]string


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestApproveUserSignInInvalidCaseLongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{"username":"dadsdsdsdfadsfasdfasdfasdfadfasdfasdfasdfaasdfasdfasdf","password":"GO"}`)

  expResponse := gin.H{"error":"Invalid username/password length"}
  var actResponse map[string]string


  w := PerformPostRequest(Router, "POST", "/api/signin", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}


func TestCreateNewUserSignUpBaseCase (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "ajay",
    "password": "GO",
    "email": "ajay@gmail.com",
    "fname": "Ajay",
    "lname": "Desai",
    "appUseStatus": "wing_mate"
    }`)

  var expDict = make(map[string]interface{})
  expDict["username"]="ajay"
	expDict["appUseStatus"]="wing_mate"
  expResponse := gin.H{"msg":expDict}


  w := PerformPostRequest(Router, "POST", "/api/signup", inputStr)
  assert.Equal(t, http.StatusOK, w.Code)

  var actResponse map[string]interface{}
  var actDict map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)
  actDict = actResponse["msg"].(map[string]interface{})

  val1, exists1 := actDict["username"]
  val2, exists2 := actDict["appUseStatus"]

  assert.Nil(t, err)
  assert.True(t, exists1)
  assert.True(t, exists2)
  assert.Equal(t, expResponse["msg"], actResponse["msg"])
  assert.Equal(t, expDict["username"], val1)
  assert.Equal(t, expDict["appUseStatus"], val2)
}

func TestCreateNewUserSignUpInvalidCaseWrongInput (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "rad5434",
    "password": "GO",
    "email": "rad5434@gmail.com",
    "fname": "Raj",
    "lname": "Desai",
    "appUseStatus": "poop"
    }`)

  expResponse := gin.H{"error":"Missing/Wrong Fields"}


  w := PerformPostRequest(Router, "POST", "/api/signup", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestCreateNewUserSignUpInvalidCaseMissingField (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "rad5434",
    "password": "GO",
    "email": "rad5434@gmail.com",
    "fname": "Raj",
    "lname": "Desai"
    }`)

  expResponse := gin.H{"error":"Missing/Wrong Fields"}


  w := PerformPostRequest(Router, "POST", "/api/signup", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestCreateNewUserSignUpInvalidCaseEmptyField (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "rad5434",
    "password": "GO",
    "email": "",
    "fname": "Raj",
    "lname": "Desai",
    "appUseStatus": "poop"
    }`)

  expResponse := gin.H{"error":"Missing/Wrong Fields"}


  w := PerformPostRequest(Router, "POST", "/api/signup", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}

func TestCreateNewUserSignUpInvalidCaseLongUsername (t *testing.T){
  m.ConnectToDB()
	defer m.DBClose()

  var inputStr = []byte(`{
    "username": "dadsdsdsdfadsfasdfasdfasdfadfasdfasdfasdfaasdfasdfasdfsdasdasd",
    "password": "GO",
    "email": "",
    "fname": "Raj",
    "lname": "Desai",
    "appUseStatus": "poop"
    }`)

  expResponse := gin.H{"error":"Missing/Wrong Fields"}


  w := PerformPostRequest(Router, "POST", "/api/signup", inputStr)
  assert.Equal(t, http.StatusBadRequest, w.Code)

  var actResponse map[string]interface{}

  err := json.Unmarshal([]byte(w.Body.String()), &actResponse)

  assert.Nil(t, err)
  assert.Equal(t, expResponse["error"], actResponse["error"])
}
