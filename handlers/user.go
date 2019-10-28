package handlers

import (
	"api/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateUser (c *gin.Context) {
	db := c.MustGet("db").(* gorm.DB)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(400)
		return
	}
	db.Create(&user)
	c.JSON(200, user)
}

func GetUser (c *gin.Context) {
	c.JSON(200, `message: "GetUser"`)
}

func GetUsers (c *gin.Context) {
	db := c.MustGet("db").(* gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(200, users)
}

func DeleteUser (c *gin.Context) {
	c.JSON(200, `message: "DeleteUser"`)
}

func PatchUser (c *gin.Context) {
	c.JSON(200, `message: "PatchUser"`)
}
