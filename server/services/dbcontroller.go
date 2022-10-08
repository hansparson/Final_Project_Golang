package services

import (
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func User_DB_Controller(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}
