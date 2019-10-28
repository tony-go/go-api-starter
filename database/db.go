package database

import (
	"api/database/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// todo : add this config to .env
const (
	host     = "localhost"
	port     = 5432
	user     = "tony_dev"
	password = "dev"
	dbName   = "goreztony"
)

// init data base => db, err := database.Init()
func Init() (*gorm.DB, error) {
	address := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	db, err := gorm.Open("postgres", address)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	migrate(db)
	fmt.Println("Database connected")
	return db, err
}

// middleware : inject DB in context => app.Use(database.Inject(db)
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// auto migrate
func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	fmt.Println("Auto migration processed !")
}
