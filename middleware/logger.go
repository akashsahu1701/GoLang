package middleware

import (
	"bytes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Process request
		c.Next()

		// Calculate response time
		latency := time.Since(start)

		// Get status code and response
		status := c.Writer.Status()

		if status >= 400 {
			//ok this is an request with error, let's make a record for it
			// now print body (or log in your preferred way)
			log.Printf("URL: %v | Response:  %v \n", c.Request.URL, blw.body.String())
		}

		// Log the request details
		log.Printf("URL: %v | Status: %d | Latency: %v", c.Request.URL, status, latency)
	}
}
