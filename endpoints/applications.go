package endpoints

import (
	"net/http"
	db "tunity-api/database/sqlite"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
)

func GetApplicationsByUserId(c *gin.Context) {
	var applications []structures.Application
	user_id := c.Param("user_id")
	err := db.GetDB().Where("user_id = ?", user_id).Find(&applications).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, applications)
}

func GetApplicationById(c *gin.Context) {
	var application structures.Application
	application_id := c.Param("id")
	err := db.GetDB().Where("id = ?", application_id).Find(&application).Limit(1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, application)
}