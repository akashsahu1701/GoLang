package main

import (
	"firstGolangModule/controllers"
	"firstGolangModule/services"

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
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "server is running on port:",
		})
	})

	server.GET("/user", func(ctx *gin.Context) {
		users, error := UserController.FindAll()
		if error != nil {
			ctx.JSON(404, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(200, users)
	})

	server.POST("/user", func(ctx *gin.Context) {
		user, error := UserController.Create(ctx)
		if error != nil {
			ctx.JSON(400, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(200, user)
	})

	server.DELETE("/user", func(ctx *gin.Context) {
		msg, error := UserController.Delete(ctx)
		if error != nil {
			ctx.JSON(404, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(200, msg)
	})

	server.PUT("/user", func(ctx *gin.Context) {
		user, error := UserController.Update(ctx)

		if error != nil {
			ctx.JSON(404, gin.H{"error": error.Error()})
			return
		}
		ctx.JSON(200, user)
	})

	server.Run(":8080")
}
