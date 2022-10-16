package main

import (
	"final/db"
	_ "final/docs"
	"final/server"
	"final/server/services"
	"os"
)

// @title My Grams APP
// @description Final Project Golang
// @version v2.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:os.Getenv("PORT")
// @contact.name Hans Parson
// @contact.email hansparson013@gmail.com

func main() {
	var PORT = os.Getenv("PORT")
	db := db.ConnectGorm()

	//// Running Service User ////
	serviceController := services.User_DB_Controller(db)
	user_service := server.UserRouther(serviceController)
	user_service.Start(":" + PORT)
	// user_service.Start(":4000")
}
