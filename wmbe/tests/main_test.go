package tests

import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
   "github.com/gin-gonic/gin"
   "github.com/stretchr/testify/assert"
   r "github.com/rajdesai5434/mah-cool-project/wmbe/routers"
   //m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
   "bytes"
   "log"
)

var Router *gin.Engine

func PerformPostRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
   req.Header.Add("Content-Type", "application/json")
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

func PerformGetRequestSingleParam(r http.Handler, method, path, key, val string) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, nil)
   q := req.URL.Query()
   q.Add(key, val)
   req.URL.RawQuery = q.Encode()
   log.Println(req.URL.String())
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

//TestMain triggers all the tests
func TestMain(t *testing.T){
  //Connect to DB
	//m.ConnectToDB()
	//defer m.DBClose()

  Router = r.SetupRouter()

}

func TestFirst(t *testing.T){
  // Build our expected body
  body := gin.H{"message": "pong",}

  // Perform a GET request with that handler.
  w := PerformGetRequestSingleParam(Router, "GET", "/api/", "", "")

  // Assert we encoded correctly,
  // the request gives a 200
  assert.Equal(t, http.StatusOK, w.Code)

  // Convert the JSON response to a map
  var response map[string]string
  err := json.Unmarshal([]byte(w.Body.String()), &response)

  // Grab the value & whether or not it exists
  value, exists := response["message"]

  // Make some assertions on the correctness of the response.
  assert.Nil(t, err)
  assert.True(t, exists)
  assert.Equal(t, body["message"], value)

}
