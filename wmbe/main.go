package main

import (
	r "github.com/rajdesai5434/mah-cool-project/wmbe/routers"
	m "github.com/rajdesai5434/mah-cool-project/wmbe/models"
)

func main() {
	//Connect to DB
	m.ConnectToDB()
	defer m.DBClose()
	
	// Set the router as the default one shipped with Gin
	router := r.SetupRouter()

	// Start and run the server
	router.Run(":5000")
}
