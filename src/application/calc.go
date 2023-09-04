package calc

import "errors"

type Calculator struct{}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func (c *Calculator) Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}
