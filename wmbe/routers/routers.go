package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rajdesai5434/mah-cool-project/wmbe/controllers"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
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
	 return router
}