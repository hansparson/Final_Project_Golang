package server

import (
	// "final/server/controllers"

	"final/server/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
	user *services.UserController
}

func UserRouther(user *services.UserController) *Router {
	return &Router{user: user}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	///// User Handlers /////////
	router.POST("/users/register", r.user.Register_User)
	// router.POST("/users/login", r.user.POST_Orders)
	// router.PUT("/users", r.user.POST_Orders)
	// router.DELETE("/users", r.user.POST_Orders)

	router.Run(port)
}
