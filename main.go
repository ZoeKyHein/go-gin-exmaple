package main

import (
	"fmt"
	"github.com/ZoeKyHein/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    10,
		WriteTimeout:   10,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
