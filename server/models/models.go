package models

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// USER DATABASE /////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////
type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age" gorm:"not null"`
	Create_At time.Time
	Update_At time.Time
}

type Photo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_Url string `json:"photo_url"`
	User_Id   int    `json:"user_id"`
	Create_At time.Time
	Update_At time.Time
}

type Comment struct {
	ID        int    `gorm:"primaryKey"`
	User_Id   int    `json:"user_id"`
	Photo_Id  int    `json:"photo_id"`
	Message   string `json:"message"`
	Create_At time.Time
	Update_At time.Time
}

type SocialMedia struct {
	ID                int    `gorm:"primaryKey"`
	Name              string `json:"name"`
	Social_Media_Url  string `json:"social_media_url"`
	Profile_Image_Url string `json:"profile_image_url"`
	User_Id           int    `json:"user_id"`
	Create_At         time.Time
	Update_At         time.Time
}

type Request_Register struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	Create_At time.Time
	Update_At time.Time
}
