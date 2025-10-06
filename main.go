package main

import (
	"log"
	"net/http"
	"strconv"
	start "tunity-api/database/sqlite"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var initialise_db = false
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

func addUser(c *gin.Context) {
	var user structures.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	creation_error := db.Create(&user).Error
	if creation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": creation_error})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

func getUserById(c *gin.Context) {
	var user structures.User

	user_id := c.Param("user_id")

	err := db.Where("id = ?", user_id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func getAllUsers(c *gin.Context) {
	var users []structures.User
	DEFAULT_SIZE := 1

	sizeStr := c.Query("size")
	size := DEFAULT_SIZE
	var err error
	if sizeStr != "" {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid size parameter"})
			return
		}
	}

	if size > 300 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "returnable user amount must not exceed 300."})
		return
	}

	result := db.Limit(size).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if int(result.RowsAffected) < size {
		c.JSON(http.StatusOK, gin.H{
			"warning": "only " + strconv.Itoa(int(result.RowsAffected)) + "/" + strconv.Itoa(size) + " users available for this request.",
			"users": users,
		})
	} else {
		c.JSON(http.StatusOK, users)
	}
}



func main() {
	if initialise_db {
		start.StartDB()
	}
	
	initDB()
	

	router := gin.Default()
	router.GET("/applications/:user_id", getApplicationsByUserId)

	router.POST("/users", addUser)
	router.GET("/users", getAllUsers)
	router.GET("/users/:user_id", getUserById)

	router.Run("localhost:8080")
}