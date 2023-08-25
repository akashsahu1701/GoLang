package controllers

import (
	"firstGolangModule/interfaces"
	"firstGolangModule/services"
	"fmt"

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
	err := ctx.ShouldBindJSON(&user)
	// validate := validator.New()
	// err := validate.Struct(user)

	if err != nil {
		return interfaces.User{}, err
	}
	return c.services.Create(user)
}

// Delete implements UserControllers.
func (c *Controller) Delete(ctx *gin.Context) (string, error) {
	userName := ctx.Param("username")
	if len(userName) < 4 {
		return "", fmt.Errorf("%v", "please enter a valid username!!!")
	}
	// err := ctx.ShouldBindJSON(&userName)
	// if err != nil {
	// 	return "", err
	// }
	return c.services.Delete(userName)
}

// FindAll implements UserControllers.
func (c *Controller) FindAll() ([]interfaces.User, error) {
	return c.services.FindAll()
}

// Update implements UserControllers.
func (c *Controller) Update(ctx *gin.Context) (interfaces.User, error) {
	var user interfaces.User
	ctx.ShouldBindJSON(&user)
	return c.services.Update(user)
}
