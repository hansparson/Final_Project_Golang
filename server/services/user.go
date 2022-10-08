package services

import (
	"encoding/json"
	"final/server/models"
	"final/server/views"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (postgres *UserController) Register_User(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	var key_data views.Request_Register

	err := json.Unmarshal([]byte(body_string), &key_data)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "REGISTER_USER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	println(key_data.Username)
	println(key_data.Email)
	println(key_data.Password)
	println(key_data.Age)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	err_create := postgres.db.Create(&models.User{
		ID:        r1.Int(),
		Username:  key_data.Username,
		Email:     key_data.Email,
		Password:  key_data.Password,
		Age:       key_data.Age,
		Create_At: time.Time{},
		Update_At: time.Time{},
	}).Error
	if err_create != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "REGISTER_USER_FAILED",
			Error:   err_create.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "REGISTER_USER_SUCCESS",
		Payload: key_data,
	})
}
