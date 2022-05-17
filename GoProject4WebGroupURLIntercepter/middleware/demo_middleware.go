package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintHello(s string) gin.HandlerFunc {

	// Middleware 是一個 type 為 gin.HandlerFunc 的函式，
	// 簽章為 func(*Context)，Context 為 gin 的環境變數用來在 middleware 間傳遞資料
	return func(c *gin.Context) {

		// before request
		fmt.Println("hello-" + s)

		// Context.Next()呼叫下一個middleware函式
		c.Next()

		// after request
		fmt.Println("bye-" + s)

	}
}
