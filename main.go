package main

import (
	"log"
	"net/http"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("tunity_test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
}


func getApplicationsByUserId(c *gin.Context) {
	var applications []structures.Application
	user_id := c.Param("user_id")
	err := db.Where("user_id = ?", user_id).Find(&applications).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, applications)
}


func main() {

	initDB()
	

	router := gin.Default()
	router.GET("/applications/:user_id", getApplicationsByUserId)
	router.Run("localhost:8080")
}