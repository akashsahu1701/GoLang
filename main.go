package main

import (
	"firstGolangModule/controllers"
	"firstGolangModule/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	userService    services.UserServices       = services.New()
	UserController controllers.UserControllers = controllers.New(userService)
)

func main() {
	// fmt.Println("hello world!!!")
	// port := 8080
	server := gin.Default()
	// server.Use((middleware.Logger()))
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "server is running on port:",
		})
	})

	server.GET("/user", func(ctx *gin.Context) {
		users, error := UserController.FindAll()
		if error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	server.POST("/user", func(ctx *gin.Context) {
		user, error := UserController.Create(ctx)
		if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	server.DELETE("/user", func(ctx *gin.Context) {
		msg, error := UserController.Delete(ctx)
		if error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, msg)
	})

	server.PUT("/user", func(ctx *gin.Context) {
		user, error := UserController.Update(ctx)

		if error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	server.Run(":8080")
}
