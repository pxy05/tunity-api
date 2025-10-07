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
