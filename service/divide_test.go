package service

import "testing"

func TestDivideSuccess(t *testing.T) {
	result, err := Divide(10, 5)
	if err != nil {
		t.Errorf("there is an error: %v", err)
	}

	if result != 2 {
		t.Errorf("the result of the division should be %f, but got %f", 2.0, result)
	}
}

func TestDivideFail(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("theres should be an error when dividing by zero")
	}
}
