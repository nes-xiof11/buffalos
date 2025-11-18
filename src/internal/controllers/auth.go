package controllers

import (
	"buffalos/src/internal/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserController) register(c *gin.Context) {
	var (
		dto domain.User
	)

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		log.Printf("error parsing request body: %s", err)
		return
	}

	ok, err := h.Service.Create(c.Request.Context(), &dto)
	if err != nil && !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user, please try again"})
		log.Printf("error creating user: %s", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": dto})
}

func (h *UserController) login(c *gin.Context) {
	var (
		dto domain.User
	)

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		log.Printf("error parsing request body: %s", err)
		return
	}

	ok, user, err := h.Service.Login(c.Request.Context(), &dto)
	if err != nil && !ok && user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if ok {
		c.JSON(http.StatusCreated, gin.H{"data": user})
	}
}

func (h *UserController) logout(c *gin.Context) {

}
