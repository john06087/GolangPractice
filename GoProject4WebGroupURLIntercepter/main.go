package main

import (
	"fmt"

	"com.web/goProject4Web/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// http://localhost:8080/demo/v1/hello
	// http://localhost:8080/demo/v1/hi
	// http://localhost:8080/demo/v2/hello
	// http://localhost:8080/demo/v2/hi

	// 1. Middleware 的用途為若多個請求路徑需要執行某些相同的邏輯時，可以透過設定 middleware 來達成。類似 Java 的 Intercepter
	// RouterGroup.Use() 設定 middleware
	router.Use(
		middleware.PrintHello("1"),
		middleware.PrintHello("2"),
		middleware.PrintHello("3"),
	)

	// 2. 用 Group 做路徑分組
	demo := router.Group("/demo")
	demo.Group("/v1").
		GET("/hello", func(c *gin.Context) {
			fmt.Println("hello-v1")
			c.JSON(200, "Hello - v1")
		}).
		GET("/hi", func(c *gin.Context) {
			fmt.Println("hi-v1")
			c.JSON(200, "Hi - v1")
		})

	demo.Group("/v2").
		GET("/hello", func(c *gin.Context) {
			fmt.Println("hello-v2")
			c.JSON(200, "Hello - v2")
		}).
		GET("/hi", func(c *gin.Context) {
			fmt.Println("hi-v2")
			c.JSON(200, "Hi - v2")
		})

	router.Run()
}
