package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 不過 http 缺少經常需要的路由處理、JSON 處理、參數驗證、中間件處理等功能（要自己刻），所以在寫網頁伺服器時通常會用Gin框架。
	// http://localhost:8080/webWithoutGin?name=GiYu
	http.HandleFunc("/webWithoutGin", func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name") // 取得 Url 上，key 為 name 的值
		content := fmt.Sprintf("hello, %s", name)
		fmt.Fprint(rw, content) // 傳去前端印出
	})

	http.ListenAndServe(":8080", nil)
}
