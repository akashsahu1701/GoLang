package main

import (
	"firstGolangModule/controllers"
	"firstGolangModule/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	userService    services.UserServices       = services.New()
	UserController controllers.UserControllers = controllers.New(userService)
)

func main() {
	// fmt.Println("hello world!!!")
	port := 8080
	server := gin.Default()
	// server.Use((middleware.Logger()))
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "server is running on port:" + strconv.Itoa(port),
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

	server.DELETE("/user/:username", func(ctx *gin.Context) {
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

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func double(num int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(2 * num)
// 	time.Sleep(1 * time.Second)
// }

// func increment(num *int, mu *sync.Mutex, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// defer mu.Unlock()
// 	// mu.Lock()
// 	*num++
// 	time.Sleep(1 * time.Millisecond)
// }

// func multipleSeven(number chan int) {
// 	wg := &sync.WaitGroup{}
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func(i int, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			if i%7 == 0 {
// 				number <- i
// 			}
// 		}(i, wg)

// 	}
// 	wg.Wait()
// 	close(number)
// }

// func main() {
// 	t := time.Now()
// 	// mu := &sync.Mutex{}
// 	// wg := &sync.WaitGroup{}
// 	// number := 0

// 	// wg.Add(10)
// 	// for i := 0; i < 10; i++ {
// 	// 	// go double(i, wg)
// 	// 	go increment(&number, mu, wg)
// 	// }
// 	// wg.Wait()
// 	// fmt.Println("number", number)
// 	// numArr := []int{}
// 	ch := make(chan int, 10)
// 	multipleSeven(ch)
// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// 	ti := time.Since(t)
// 	fmt.Println("time taken::::::::::::: ", ti)

// }
