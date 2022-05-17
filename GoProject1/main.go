package main // 宣告此程式的 package 名稱為 main (執行檔必須在 main 下)

import (
	"errors"
	"fmt" // 匯入標準函式庫的fmt package，其提供文字格式化與console列印等函式
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	greetings "abc.com/GoProject2" // 引入外部專案
	"go.uber.org/zap"
	"rsc.io/quote" // 匯入外部package rsc.io/quote
	//import "github.com/gin-gonic/gin"	// 沒用到的要註解，不然會噴錯
)

const (
	constNum1 = iota
	constNum2 = iota
	constNum3 = iota
)

func main() { // 定義 main 函式，其為執行 main package 時預設會執行的函式，即程式進入點

	fmt.Println("範例 1 - print hello world, 引入開源專案")
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go()) // 呼叫 rsc.io/quote 的 Go() 函式

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 2 - 各種宣告方式")
	var array1 [3]int = [3]int{1, 2, 3}
	array2 := [5]int{1: 1, 2: 2, 3: 3}
	var b string = "Hi~~~"
	var c func(string, string) string = func(x string, y string) string { return "AP" + "ple" }
	var d int               // 預設值 0
	var e bool              // 預設值 false
	var f = true            // 自動判斷型別
	var g, h int = 100, 200 // 一次宣告兩個
	i := "hello 什麼都不用宣告"

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 3 - 引入本地專案")
	message := greetings.Hello("Lin-Meng")
	fmt.Println(message)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 4 - 指標")
	testCase := 1 // 不可使用布林
	var j *int    // 加上 "*" 代表此變數儲存指標位置
	j = &testCase // 加上 "&" 取得指標位置
	fmt.Println(j)

	fmt.Println(*j) // 指標位置加上 "*" 代表設定該指標的值
	var changeIn = *j
	changeIn = 2
	fmt.Println(changeIn) // 拿出來儲存到另一個變數後就不會和 testCase 同步了
	*j = 90
	fmt.Println(testCase) // 90

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 5 - Slice, Map, Delete, Foreach")
	// Map
	holiday := make(map[string]string)
	holiday["1225"] = "聖誕節"
	holiday["0815"] = "中秋節"
	holiday["0808"] = "父親節"

	val, exist := holiday["1225"] // 檢查是否存在
	if exist {
		fmt.Println("value is " + val)
	} else {
		fmt.Println("no value")
	}
	fmt.Println(len(holiday)) // 查看長度
	fmt.Println(holiday)
	delete(holiday, "1225")
	fmt.Println(holiday)

	// Slice
	sliceTest := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sliceTest); i++ {
		fmt.Println(sliceTest[i])
	}

	// Slice delete
	sliceTest = append(sliceTest[:3], sliceTest[4:]...) // 刪除第三個元素...

	// forEach
	for index, value := range sliceTest {
		fmt.Printf("[%d]: %d\n", index, value)
	}
	for key, value := range holiday {
		fmt.Printf("[%s]: %s\n", key, value)
	}

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 6 - error handler")
	err := errors.New("empty name")
	if err == nil {
		log.Fatal(err) // 列印錯誤並中斷程式
	}

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 7 - function")
	firstName, secondName := testSplitNameFunc("Grimmjow-Jaegerjaques")
	fmt.Println(firstName)
	fmt.Println(secondName)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 8 - switchCase") // 與 Java 不同，不需要 break
	rand.Seed(time.Now().UnixNano())
	xNumber := rand.Intn(10)
	switch xNumber {
	case 1, 2, 3:
		fmt.Println("is 1 2 3")
	case 4, 5, 6, 7:
		fmt.Println("is 4 5 6 7")
	default:
		fmt.Println("is 8 9 10")
	}

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 9 - struct") // 類似 Java 的 Bean
	employee1 := Employee{1, "Dora", 27}
	employee2 := Employee{id: 2, name: "BCW"}
	fmt.Println(employee1)
	fmt.Println(employee2)
	employee2.age = 3
	fmt.Println(employee2)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 10 - Map 簡單操作")
	testMap := make(map[int]string)
	testMap[1] = "one"
	testMap[10] = "ten"
	fmt.Println("testMap[1]: " + testMap[1] + ", testMap[2]: " + testMap[2] + ", testMap[10]" + testMap[10])
	v, exist := testMap[1]
	fmt.Println(v)
	fmt.Println(exist)
	v, exist = testMap[2]
	fmt.Println(v)
	fmt.Println(exist)
	for k, v := range testMap {
		fmt.Printf("key=%d, value=%s\n", k, v)
	}

	testMap2 := map[int]string{
		3: "three",
		6: "six",
	}
	fmt.Println(testMap2)
	delete(testMap2, 3)
	fmt.Println(testMap2)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 11 - Slice 一些方法 len cap append copy")
	numberB := []int{-5, -4, -3, -2, -1}
	numberC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy(numberC, numberB)
	fmt.Print("numberB: ")
	fmt.Println(numberB)
	fmt.Print("numberC: ")
	fmt.Println(numberC)
	numberB = append(numberB, 0)
	numberB = append(numberB, 1)
	fmt.Print("numberB: ")
	fmt.Println(numberB)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("測驗 12 - 回傳正整數")
	examSlice := []int{-2, -1, 0, 4, 5}
	answerSlice := returnPostive(examSlice)
	fmt.Println(answerSlice)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 13 - Receiver")

	receiverEmployee := Employee{99, "Sin", 29}
	receiverEmployee.fakeAddAge()
	fmt.Println("fakeAddAge: " + strconv.Itoa(receiverEmployee.age))
	receiverEmployee.realAddAge()
	fmt.Println("realAddAge: " + strconv.Itoa(receiverEmployee.age))

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 14 - Receiver")
	var applefruit Fruit = Apple{"red", "circle", "sweet"}
	fmt.Println(applefruit.tasteLike())
	// interface 有幾個限制
	// 1. Interface 方法內可以取得實踐 struts 的值，Ex: tasteLike 內可以 fruit(apple) 的值
	//	  但方法外取不到 fmt.Println(applefruit.color) <- 會失敗
	// 2. 必須實現 interface 裡面的所有方法才算實現了 interface
	// 3. interface 不允許相同方法名稱 Ex: tasteLike() + tasteLike(s string)
	// 4. 兩個介面可以融合在一起

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 15 - 讀取文字檔")
	loadTxtByte, loadTxterr := ioutil.ReadFile("讀取.txt")
	if loadTxterr != nil {
		log.Fatal(loadTxterr)
	}
	loadTxtStr := string(loadTxtByte) // string 將數字转换成 ASCII 碼同等的字符  vs  strconv.Itoa
	fmt.Println(loadTxtStr)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 16 - 暫停目前的 goroutine")
	fmt.Println("first")
	time.Sleep(time.Second * 1) // 1000000000
	fmt.Println("second")

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 17 - goroutine 併行程序 (Java MultiThread)")
	// go count(3, "Goroutine: ") //	兩條 Thread 並行處理
	// count(3, "Main: ")

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 18 - 使用 Channel 在 goroutine 間傳遞資料")
	// 注意: 如果 set 值進 channel 沒有拿出來，該 Thread 會停止，並且 1 個 set 會搭配 1 個 get 執行，全部 set 執行完畢才會往下做
	//		 channel 的收發要同時可以執行才會成功，否則發生 deadlock，解決方式可新增 channel 的 buffer
	channelTest1 := make(chan string)
	go send("hello", channelTest1)
	stringTest1 := <-channelTest1
	fmt.Println(stringTest1)
	// 先進先出，達到最大值才會 deadlock，deadlock 的如果是 main Thread 會噴掉
	channelTest2 := make(chan string, 3)
	channelTest2 <- "hello1"
	channelTest2 <- "hello2"
	fmt.Println(<-channelTest2)
	fmt.Println(<-channelTest2)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 19 - struct 嵌入欄位 embedded field")
	// Go 並沒有物件導向程式如 Java 的繼承設計，而是透過嵌入欄位來達到類似效果。所以上面的 Manager 就像是"繼承"了 Employee 的特性。
	var fob FakeObj = FakeObj{90, "fakeObjjjjjj"}
	fmt.Println("fakeObj: " + fob.string)
	var persianCat Animal = Animal{Cat{5, "persian", true}, "burberry"}
	fmt.Println("貓的種類: " + persianCat.catType)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 20 - Golang 在for loop修改struct元素的屬性")
	empsForChangeValue := []Employee{
		{id: 1, name: "John", age: 33},
		{id: 1, name: "Mary", age: 28},
	}
	for _, empForChange := range empsForChangeValue {
		empForChange.age = empForChange.age + 1
	}
	fmt.Print("年齡並未改變: ")
	fmt.Println(empsForChangeValue)

	for i := 0; i < len(empsForChangeValue); i++ {
		emp := &empsForChangeValue[i] // get pointer of emps[i]
		emp.age = emp.age + 100
	}
	fmt.Print("年齡跟著改變: ")
	fmt.Println(empsForChangeValue)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 21 - Continue 回到指定的 for 迴圈")
	depSlice := []Department{
		{
			id:   1,
			name: "HR",
			employees: []Employee{
				{id: 1, name: "John", age: 33},
				{id: 2, name: "Mary", age: 28},
			},
		},
		{
			id:   2,
			name: "Mkt.",
			employees: []Employee{
				{id: 3, name: "Bill", age: 18},
				{id: 4, name: "Nick", age: 17},
			},
		},
	}
	ageLess30Result := filterDeptAllEmpAgeLessThen30(depSlice)
	fmt.Println(ageLess30Result)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 22 - iota 只能用在常數")
	fmt.Println(constNum1)
	fmt.Println(constNum2)
	fmt.Println(constNum3)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 23 - Enum")
	// Enum 範例 1
	fmt.Printf("%d, %d, %d\n", Black, Red, Blue)                            // 0, 1, 2
	fmt.Printf("%s, %s, %s\n", Black.String(), Red.String(), Blue.String()) // Black, Red, Blue
	colorNum := Red
	switch colorNum {
	case Black:
		fmt.Println("黑色")
	case Red:
		fmt.Println("紅色")
	case Blue:
		fmt.Println("藍色")
	}

	// Enum 範例 2
	resBk, err := Of("822")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resBk.Name())

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 24 - Trim 範例")
	testTrim := " hello  "

	trimResult1 := strings.Trim(testTrim, " ")
	trimResult2 := strings.TrimLeft(testTrim, " ")
	trimResult3 := strings.TrimRight(testTrim, " ")
	trimResult4 := strings.TrimSpace(testTrim)
	trimResult5 := strings.TrimPrefix(testTrim, " he")
	trimResult6 := strings.TrimSuffix(testTrim, "llo  ")
	fmt.Println("{" + trimResult1 + "}")
	fmt.Println("{" + trimResult2 + "}")
	fmt.Println("{" + trimResult3 + "}")
	fmt.Println("{" + trimResult4 + "}")
	fmt.Println("{" + trimResult5 + "}")
	fmt.Println("{" + trimResult6 + "}")

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 25 - Build exe - go build example.com/hello ")

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 26 - uber go zap log 簡單範例 - 下載: go get -u go.uber.org/zap")
	// NewExample()、NewDevelopment()、NewProduction()
	// 差別參考以下: https://matthung0807.blogspot.com/2021/09/go-uber-go-zap-presets-logger.html
	sugar := zap.NewExample().Sugar() // 效能較差，但提供訊息格式化 API
	defer sugar.Sync()

	sugar.Debug("debug message")
	sugar.Info("info message")
	sugar.Error("error message")
	sugar.Warn("warn message")
	// sugar.Panic("panic message")
	// sugar.Fatal("fatal message")

	logger := zap.NewExample() // 效能較好
	defer logger.Sync()

	logger.Info("failed to fetch URL",
		zap.String("url", "http://abc.com/get"),
		zap.Int("attempt", 3),
		zap.Bool("enabled", true))

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 27 - 產生隨機整數 generate int random number ")

	fmt.Println(rand.Intn(10)) // 不重新執行的話，每次產生的結果都會相同
	// ---------------------------
	rand.Seed(time.Now().UnixNano()) // 搭配種子使用產生不同的亂數
	fmt.Println(rand.Intn(10))
	// ---------------------------
	minRand := 10
	maxRand := 20
	n := rand.Intn(maxRand-minRand) + minRand // 產生 10 ~ 20 的亂數
	fmt.Println(n)

	fmt.Println("======== ======== ======== ======== ======== ======== ========")

	fmt.Println("範例 27 - 時間 ")
	nowTime := time.Now()
	fmt.Println(nowTime.Format("20060102"))

}

type FakeObj struct {
	int
	string
}

type Department struct {
	id        int
	name      string
	employees []Employee
}

type Employee struct {
	id   int
	name string
	age  int
}

type Animal struct {
	Cat
	collar string // 項圈
}
type Cat struct {
	age        int
	catType    string
	isligation bool // 是否結紮
}

type Fruit interface {
	tasteLike() string
}

type Apple struct {
	color  string
	shape  string
	flavor string
}

func testSplitNameFunc(realName string) (string, string) {
	realNameSplit := strings.Split(realName, "-")

	firstName := realNameSplit[0]
	secondName := realNameSplit[1]

	return firstName, secondName
}

func (e Employee) fakeAddAge() {
	e.age += 100
	fmt.Println("fakeAddAge: " + strconv.Itoa(e.age))
}

func (e *Employee) realAddAge() {
	e.age += 200
	fmt.Println("realAddAge: " + strconv.Itoa(e.age))
}

func returnPostive(checkSlice []int) []int {
	resultSlice := []int{}

	for _, v := range checkSlice {
		if v > 0 {
			resultSlice = append(resultSlice, v)
		}
	}

	return resultSlice
}

func (a Apple) tasteLike() string {
	return "The fruit taste like: " + a.flavor
}

func count(n int, s string) {
	for i := 0; i <= n; i++ {
		fmt.Println(s + strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
}
func send(s string, c chan string) {
	c <- s // send s to channel c
}

func filterDeptAllEmpAgeLessThen30(depSlice []Department) []Department {
	var result []Department

OuterFor:
	for _, dep := range depSlice {
		allEmpLess30 := true //是否部門內全部員工歲數 < 30
		empSlice := dep.employees
		for _, emp := range empSlice {
			if allEmpLess30 && emp.age > 30 {
				allEmpLess30 = false
				continue OuterFor
			}
		}

		result = append(result, dep)
	}

	return result
}
