package endpoints

import (
	"encoding/json"
	"net/http"
	supabase "tunity-api/database/supabase"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
)

func SBGetApplicationsByUserID(c *gin.Context) {
	db := supabase.GetDB()
	userID := c.Param("user_id")

	data, _, err := db.From("applications").Select("*", "exact", false).Filter("user_id", "eq", userID).Execute()

	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	applications := []structures.Application{}
	if err := json.Unmarshal([]byte(string(data)), &applications); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, applications)
}

func SBAddApplication(c *gin.Context) {
	db := supabase.GetDB()
	var application structures.Application

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if application.UserID == nil || *application.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if application.Company == nil || *application.Company == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company is required"})
		return
	}

	_, _, err := db.From("applications").Insert(application, false, "", "", "").Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, application)
}