package server

import (
	// "final/server/controllers"

	"final/server/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	control *services.HandlersController
}

func UserRouther(control *services.HandlersController) *Router {
	return &Router{control: control}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	///// User Handlers /////////
	router.POST("/users/register", r.control.Register_User)
	router.POST("/users/login", r.control.Login_User)
	router.PUT("/users", r.control.PUT_User)
	router.DELETE("/users", r.control.Delete_User)

	router.POST("/photos", r.control.Post_Photos)
	router.GET("/photos", r.control.Get_Photos)
	router.PUT("/photos/:photoId", r.control.Put_Photos)
	router.DELETE("/photos/:photoId", r.control.Delete_Photos)

	router.POST("/comments", r.control.Comments_Post)
	router.GET("/comments", r.control.Comment_Get)
	router.PUT("/comments/:commentId", r.control.Comment_Put)
	router.DELETE("/comments/:commentId", r.control.Comment_Delete)

	router.POST("/socialmedias", r.control.SocialMedias_Post)
	router.GET("/socialmedias", r.control.SocialMedias_Get)
	router.PUT("/socialmedias/:socialMediaId", r.control.SocialMedias_Put)
	router.DELETE("/socialmedias/:socialMediaId", r.control.SocialMedias_Delete)

	router.Run(port)
}
