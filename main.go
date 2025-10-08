package main

import (
	supabase "tunity-api/database/supabase"
	"tunity-api/endpoints"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)





func main() {

	
	godotenv.Load()
	supabase.InitDB(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_ROLE_KEY"))
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/users/applications/:user_id", endpoints.SBGetApplicationsByUserID)
	router.POST("/applications", endpoints.SBAddApplication)
	// router.GET("/applications/:id", endpoints.GetApplicationById)

	// router.POST("/users", endpoints.AddUser)
	router.GET("/users", endpoints.SBGetAllUsers)
	//router.GET("/users/:user_id", endpoints.SBGetUserById)

	router.Run("localhost:8080")
}