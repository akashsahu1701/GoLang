package controllers

import (
	"firstGolangModule/interfaces"
	"firstGolangModule/services"

	"github.com/gin-gonic/gin"
)

type UserControllers interface {
	Create(ctx *gin.Context) (interfaces.User, error)
	Update(ctx *gin.Context) (interfaces.User, error)
	FindAll() ([]interfaces.User, error)
	Delete(ctx *gin.Context) (string, error)
}

type Controller struct {
	services services.UserServices
}

func New(services services.UserServices) UserControllers {
	return &Controller{
		services: services,
	}
}

// Create implements UserControllers.
func (c *Controller) Create(ctx *gin.Context) (interfaces.User, error) {
	var user interfaces.User
	ctx.BindJSON(&user)
	return c.services.Create(user)
}

// Delete implements UserControllers.
func (c *Controller) Delete(ctx *gin.Context) (string, error) {
	var username string
	ctx.BindJSON(&username)
	return c.services.Delete(username)
}

// FindAll implements UserControllers.
func (c *Controller) FindAll() ([]interfaces.User, error) {
	return c.services.FindAll()
}

// Update implements UserControllers.
func (c *Controller) Update(ctx *gin.Context) (interfaces.User, error) {
	var user interfaces.User
	ctx.BindJSON(&user)
	return c.services.Update(user)
}
