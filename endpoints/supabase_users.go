package endpoints

import (
	"encoding/json"
	"net/http"
	supabase "tunity-api/database/supabase"
	"tunity-api/tools/structures"

	"github.com/gin-gonic/gin"
)

func SBGetAllUsers(c *gin.Context) {

	just_id := c.Query("just_id") == "true"

	select_string := "*"
	if just_id {
		select_string = "id"
	}

	db := supabase.GetDB()
	data, _, err := db.From("users").Select(select_string, "", false).Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users := []structures.User{}
	if err := json.Unmarshal([]byte(string(data)), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func SBAddUser(c *gin.Context) {
	db := supabase.GetDB()
	user := structures.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID == nil || *user.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	_, _, err := db.From("users").Insert(user, false, "", "", "").Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func SBUpdateUser(c *gin.Context) {
	db := supabase.GetDB()
	user := structures.User{}
	user_id := c.Param("user_id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user.ID = &user_id

	if user.ID == nil || *user.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	_, _, err := db.From("users").Update(user, "", "").Filter("id", "eq", user_id).Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
