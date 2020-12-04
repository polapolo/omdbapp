package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// InitHTTPServer -> return http server to serve
func InitHTTPServer() *http.Server {
	// routes
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// server
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
