package controllers

import (
	"buffalos/src/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.User
}

func NewUserController(service *services.User) *UserController {
	return &UserController{Service: service}
}

func (h *UserController) RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", h.list)
		users.GET("/:id", h.get)
		users.PUT("/:id", h.update)
		users.DELETE("/:id", h.delete)
	}

	auth := router.Group("/u")
	{
		auth.POST("/", h.register)
		auth.POST("/login", h.login)
		auth.POST("/logout", h.logout)
	}
}

func (h *UserController) get(c *gin.Context) {

}

func (h *UserController) list(c *gin.Context) {

}

func (h *UserController) update(c *gin.Context) {

}

func (h *UserController) delete(c *gin.Context) {

}
