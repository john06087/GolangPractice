package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "com.web/GoProject5Web/docs"
)

// 以下註解通通必要，不可刪除

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
func main() {

	router := gin.Default()

	router.Group("/demo/v1").
		GET("/hello", hello).
		GET("/hi", hi)

	// http://localhost:8080/swagger/index.html
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}

// @Success 200 {string} string
// @Router /demo/v1/hello [get]
func hello(c *gin.Context) {
	c.JSON(200, "Hello - v1")
}

// @Success 200 {string} string
// @Router /demo/v1/hi [get]
func hi(c *gin.Context) {
	c.JSON(200, "Hi - v1")
}
