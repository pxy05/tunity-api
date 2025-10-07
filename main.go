package main

import (
	db "tunity-api/database/sqlite"
	"tunity-api/endpoints"

	"github.com/gin-gonic/gin"
)

var initialise_db = true


func main() {
	if initialise_db {
		db.StartDB()
	}
	
	router := gin.Default()
	router.GET("/users/applications/:user_id", endpoints.GetApplicationsByUserId)
	router.GET("/applications/:id", endpoints.GetApplicationById)

	router.POST("/users", endpoints.AddUser)
	router.GET("/users", endpoints.GetAllUsers)
	router.GET("/users/:user_id", endpoints.GetUserById)

	router.Run("localhost:8080")
}