package tests

import(
  "encoding/json"
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  r "github.com/rajdesai5434/mah-cool-project/wmbe/routers"
)

func TestApproveUserSignIn (t *testing.T){
  reqBody := gin.H{"username": "rad5434","password":"GO",}
  rightResponse := gin.H{"username": "rad5434","appUseStatus":"wing_mate",}

  router := r.SetupRouter()
  w := PerformPostRequest(router, "POST", "/api/signin", reqBody)
  assert.Equal(t, http.StatusOK, w.Code)

  var reqResponse map[string]string
  err := json.Unmarshal([]byte(w.Body.String()), &reqResponse)

  val1, exists1 := reqResponse["username"]
  val2, exists2 := reqResponse["appUseStatus"]

  assert.Nil(t, err)
  assert.True(t, exists1)
  assert.True(t, exists2)
  assert.Equal(t, rightResponse["username"], val1)
  assert.Equal(t, rightResponse["appUseStatus"], val2)
}
