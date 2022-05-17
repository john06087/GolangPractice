package TestService

import (
	"errors"
	Model "go/Model"
)

type CalculatorService interface {
	Plus(int, int) int
}

type Calculator struct {
}

func (c Calculator) Plus(x, y int) int {
	return x + y
}

func AddAge(x int, emp Model.Employee, calService CalculatorService) (int, error) {
	if (emp == Model.Employee{}) {
		return -1, errors.New("emp is empty")
	}
	return calService.Plus(x, emp.Age), nil
}
