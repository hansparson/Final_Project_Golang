package db

import (
	"final/server/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// LOCAL DB
// const (
// 	host   = "localhost"
// 	port   = "5432"
// 	user   = "postgres"
// 	pass   = "12345"
// 	dbname = "my_gram"
// )

// Railway DB
const (
	host   = "containers-us-west-49.railway.app"
	port   = "7132"
	user   = "postgres"
	pass   = "UhTOywF4FVz3TTRddYmW"
	dbname = "Postgres"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		// panic(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect to database")
	}

	// db.Debug().AutoMigrate(models.Item{})
	// db.Debug().AutoMigrate(models.Orders{})

	/// Migrate Mygram Database
	errors := db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}
