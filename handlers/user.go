package handlers

import "github.com/gin-gonic/gin"

func CreateUser (c *gin.Context) {
	c.JSON(200, `message: "CreateUser"`)
}

func GetUser (c *gin.Context) {
	c.JSON(200, `message: "GetUser"`)
}

func GetUsers (c *gin.Context) {
	c.JSON(200, `message: "GetUsers"`)
}

func DeleteUser (c *gin.Context) {
	c.JSON(200, `message: "DeleteUser"`)
}

func PatchUser (c *gin.Context) {
	c.JSON(200, `message: "PatchUser"`)
}
