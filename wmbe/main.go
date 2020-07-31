package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajdesai5434/mah-cool-project/wmbe/controllers"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)


func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	m.ConnectToDB()
	defer m.DBClose()

	//sign.go
	api.POST("/signup", controllers.CreateNewUserPost)
	api.POST("/signin",controllers.ApproveUserSignIn)

	//profile.go
	api.GET("/wingmate_profile/:username",controllers.GetWingmateProfile)
	api.POST("/wingmate_profile",controllers.PostWingmateProfile)
	api.GET("/dater_profile/:username",controllers.GetDaterProfile)
	api.POST("/dater_profile",controllers.PostDaterProfile)


	//api.POST("/signin",controllers.ApproveUserSignIn)

	// Our API will consit of just two routes
	// /jokes - which will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - which will capture likes sent to a particular joke
	//api.GET("/jokes", JokeHandler)
	//api.POST("/jokes/like/:jokeID", LikeJoke)

	// Start and run the server
	router.Run(":5000")
}

/*
// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, jokes)
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
  // confirm Joke ID sent is valid
  // remember to import the `strconv` package
  if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
    // find joke, and increment likes
    for i := 0; i < len(jokes); i++ {
      if jokes[i].ID == jokeid {
        jokes[i].Likes += 1
      }
    }

    // return a pointer to the updated jokes list
    c.JSON(http.StatusOK, &jokes)
  } else {
    // Joke ID is invalid
    c.AbortWithStatus(http.StatusNotFound)
  }
}
*/
