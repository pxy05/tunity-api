package endpoints

import (
	"encoding/json"
	"net/http"
	supabase "tunity-api/database/supabase"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
)


func SBEditApplication(c *gin.Context) {
	db := supabase.GetDB()
	appli_id := c.Param("appli_id")
	var updated_appli structures.Application
	
	if err := c.ShouldBindJSON(&updated_appli); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	
	if updated_appli.UserID != nil && *updated_appli.UserID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot change user_id for an application"})
		return
	}
	
	if updated_appli.ID != nil && *updated_appli.ID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot change application's ID for an application"})
		return
	}

	updated_appli.ID = &appli_id
	
	data, _, err := db.From("applications").Update(updated_appli, "", "").Filter("id", "eq", appli_id).Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []structures.Application
	if err := json.Unmarshal(data, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
	
}

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
