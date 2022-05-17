package GoProject2

import (
	"testing"

	Model "go/Model"
	"goTest/TestService"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 注意事項
// 1. 黨案名稱需要以 _test 結尾，golang 才能偵測到
// 2. 每支測試函式的名稱以Test開頭，後接被測函式的名稱，固定傳入(t *testing.T)
// 3. Go雖然自帶 testing package 可撰寫測試程式但功能陽春，沒有常見的 assert 或 mock 功能，而 Testify 套件則彌補這塊不足讓撰寫測試更方便
//    在專案根目錄以命令列輸入 go get github.com/stretchr/testify 下載 Testify module 及依賴
//	  Testify功能包括：
//	  	assert：提供方便的測試方法，省去寫if-else判斷及增加可讀性。
//	 	mock：測試對象的依賴物件可用假的物件(test double)來替代達到單元測試的隔離效果
//		suite：組織多個測試案例為測試包/測試套組（test suite）
// 4. 清除測試快取 go clean -testcache
// 5. 在專案根目錄命令列輸入go test -v ./...即可執行專案中所有目錄中的測試程式，./代表相對路徑；/...代表任意字串

func TestPlusSimple(t *testing.T) {
	// Arrange - prepare test cases, input arguments and expected result.
	testCases := []struct {
		c, w, x, y, z, expected int
	}{
		{1, 2, 3, 4, 5, 5},
		{6, 7, 8, 9, 10, 20},
	}

	for _, testCase := range testCases {
		// Act - call function and pass test case arguments
		result := Plus(testCase.c, testCase.w, testCase.x, testCase.y, testCase.z)

		// Assert - compare actual result with expected result
		if result != testCase.expected {
			t.Errorf("錯了錯了")
		}
	}
}

//---------------------------------------------------------------

func TestMinusAssert(t *testing.T) {
	testCases := []struct {
		x, y, expected int
	}{
		{2, 1, 1},    // test case 1
		{3, 1, 2},    // test case 2
		{100, 1, 99}, // test case 3
	}

	for _, testCase := range testCases { // execute test cases one by one
		result := Minus(testCase.x, testCase.y)
		assert.Equal(t, result, testCase.expected)
	}
}

//---------------------------------------------------------------

func TestAddAgeMock(t *testing.T) {

	// 建立 Mock 的 Plus 方法，傳入 1 & 33 回傳 100 (正常是 34)
	calMock := new(CalculatorMock)
	calMock.On("Plus", 1, 33).Return(100)

	// 傳入 John 33 + 1 歲
	emp2 := Model.Employee{Id: 1, Name: "John", Age: 33}
	actual, _ := TestService.AddAge(1, emp2, calMock) // pass test arguments and mock to replace real one

	assert.Equal(t, 100, actual)
}

// define mock type
type CalculatorMock struct {
	mock.Mock
}

// use mock to implments CalculatorService's method
func (calMock *CalculatorMock) Plus(x, y int) int {
	args := calMock.Called(x, y)
	return args.Int(0)
}
