package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "abc.com/demo/docs"
)

func main() {
	gin.DisableConsoleColor()
	// 建立 Log
	file, _ := os.Create("firstWebLog.log")             // create log file
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // wirte log to file and console
	router := gin.Default()                             // get gin engine with Logger attached

	// 1. get 取得 param, 網址列
	// URL: http://localhost:8081/hello/%E7%BE%8E%E8%B3%BD%E5%93%A6?paramA=%E3%84%8F%E3%84%8F%E3%84%8F
	router.GET("/hello/:pathParam", func(ginContext *gin.Context) {
		pathParam := ginContext.Param("pathParam") // 取得路徑參數
		paramA := ginContext.Query("paramA")
		paramB := ginContext.DefaultQuery("paramB", "default b la")

		ginContext.JSON(200, gin.H{ // response json
			"message":   "hello world",
			"pathParam": pathParam,
			"paramA":    paramA,
			"paramB":    paramB,
		})
	})

	// 2. post 取得 form 參數
	// testForm.html
	router.POST("/formTest", func(ginContext *gin.Context) {
		formInput1 := ginContext.PostForm("formInput1")
		formInput2 := ginContext.DefaultPostForm("formInput2", "default post form la")

		ginContext.JSON(200, gin.H{ // response json
			"formInput1": formInput1,
			"formInput2": formInput2,
		})
	})

	// 3. get 取得 map 類型參數
	// URL http://localhost:8081/getMap?fruitMap[1]=apple&fruitMap[2]=banana&fruitMap[3]=lemon
	router.GET("/getMap", func(ginContext *gin.Context) {
		fruitMap := ginContext.QueryMap("fruitMap") // 取得 Map 參數

		ginContext.JSON(200, gin.H{ // response json
			"fruitMap": fruitMap,
		})
	})

	// 4. post 讀取單個上傳檔案
	// testUploadFile.html
	router.POST("/uploadFile", func(ginContext *gin.Context) {
		file, _ := ginContext.FormFile("fileInput") // get file from form input name 'file'

		ginContext.SaveUploadedFile(file, "tmp/"+file.Filename) // save file to tmp folder in current directory
		ginContext.String(http.StatusOK, "File name: %s, File size: %d %s, File Header %s", file.Filename, file.Size, "byte", file.Header)
	})

	// 5. post 讀取多個上傳檔案
	// testUploadMultipleFile.html
	router.POST("/uploadMultipleFile", func(ginContext *gin.Context) {
		form, _ := ginContext.MultipartForm()
		multipleFiles := form.File["multipleFiles"]

		for _, file := range multipleFiles {
			log.Println(file.Filename)

			ginContext.SaveUploadedFile(file, "tmp/"+file.Filename)
		}
		ginContext.String(http.StatusOK, "%d files uploaded!", len(multipleFiles))
	})

	// 6. binding json to type (Resful API)
	// localhost:8081/getEmployeeJsonObj
	// {"id": 1, "name": "john", "age": 33}
	router.POST("/getEmployeeJsonObj", func(ginContext *gin.Context) {
		var emp Employee
		err := ginContext.ShouldBindJSON(&emp) // pass as pointer
		if err != nil {
			ginContext.String(http.StatusBadRequest, "error")
			return
		}

		log.Println(emp)
		ginContext.String(http.StatusOK, emp.String())
	})

	// 7. 讀取靜態檔
	// http://localhost:8081/testJpg1
	// http://localhost:8081/testJpg2/test2.jpg
	// http://localhost:8081/testJpg3
	router.StaticFile("/testJpg1", "./resources/test1.jpg") //設定單一檔案的資源位址
	router.Static("/testJpg2", "./static")                  //設定目錄為靜態資源路徑
	router.GET("/testJpg3", func(c *gin.Context) {
		c.File("./resources/test1.jpg")
	})

	// 8. Gin 轉 Request body JSON 為 struct
	// localhost:8081/cakes - Post
	// {"id":100,"name":"HR","cakes":[{"id":1,"name":"coffee cake","count":33},{"id":2,"name":"Tiramisu","count":28}]}
	router.POST("/cakes", func(ginContext *gin.Context) {
		var dessert Dessert
		err := ginContext.Bind(&dessert)
		if err != nil {
			return
		}
		fmt.Printf("%v\n", dessert)
		ginContext.JSON(200, "success")
	})

	// 9. Gin 轉 Request body JSON 為 map
	// localhost:8081/cakesToMap - Post
	// {"id":100,"name":"HR","cakes":[{"id":1,"name":"coffee cake","count":33},{"id":2,"name":"Tiramisu","count":28}]}
	router.POST("/cakesToMap", func(ginContext *gin.Context) {
		var m map[string]interface{}
		err := ginContext.Bind(&m)
		if err != nil {
			return
		}

		fmt.Printf("%v\n", m)

		ginContext.JSON(200, "success")
	})

	router.Run(":8081") // 原本 8080 改成 8081
}

type Employee struct {
	Id   int
	Name string `binding:"required"` // 代表此蘭未必填
	Age  int
}

type Dessert struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Cakes []Cake `json:"cakes"`
}

type Cake struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (e Employee) String() string {
	return "Employee={" +
		"Id=" + strconv.Itoa(e.Id) + "," +
		"Name=" + e.Name + "," +
		"Age=" + strconv.Itoa(e.Age) +
		"}"
}
