package main

import (
	"api/database"
	"api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main () {

	// init db
	db, _ := database.Init()

	// port
	port := "2910"

	// create gin app
	app := gin.Default()

	// middleware
	app.Use(database.Inject(db))

	// routes
	routes.ApplyRoutes(app)

	// app starter
	err := app.Run(":" + port)

	if err != nil {
		panic(err)
	}
}
