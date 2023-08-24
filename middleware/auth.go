package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"akashsahu1701": "akash@1111",
		"pankaj":        "akash@1111",
	})
}
