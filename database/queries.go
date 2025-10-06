package sqlite

//endpoints from db to api

import (
	"tunity-api/tools/structures"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)



getUserApplicationsByUserId(id int) (structures.Application, error) {

	var applications structures.Application[]
	err := db.Where("user_id = ?", id).Find(&applications).Error
	return applications, err

}
