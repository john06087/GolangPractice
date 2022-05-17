package GoProject2

import "errors"

func Hello(name string) string {
	fullname := name + "Shuan"
	return fullname
}

func ReturnError() (string, error) {
	error := errors.New("empty name")
	return "測試錯誤", error
}
