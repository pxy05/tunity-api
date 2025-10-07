package endpoints

import (
	"net/http"
	"strconv"
	db "tunity-api/database/sqlite"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
)


func AddUser(c *gin.Context) {
	var user structures.User
	err := c.ShouldBindJSON(&user)
	creation_error := db.GetDB().Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if creation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": creation_error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}


func GetUserById(c *gin.Context) {
	var user structures.User

	user_id := c.Param("user_id")

	err := db.GetDB().Where("id = ?", user_id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
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

	result := db.GetDB().Limit(size).Find(&users)
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
