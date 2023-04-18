package service

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot perform division by zero")
	}
	result := a / b
	return result, nil
}
