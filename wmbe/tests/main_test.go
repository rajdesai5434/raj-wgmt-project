package tests

import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
   "github.com/gin-gonic/gin"
   "github.com/stretchr/testify/assert"
   "io"
)

func PerformPostRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, body)
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

func PerformGetRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, nil)
   w := httptest.NewRecorder()
   r.ServeHTTP(w, req)
   return w
}

func TestMain(t *testing.T){
  // Build our expected body
  body := gin.H{"message": "pong",}

  router := SetupRouter()

  // Perform a GET request with that handler.
  w := PerformGetRequest(router, "GET", "/api/")

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
