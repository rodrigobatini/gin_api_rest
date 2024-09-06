package migrations

import (
	"gin_api_rest/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
