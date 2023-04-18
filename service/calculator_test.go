package service

import (
	"log"
	"testing"
)

func TestCalculatorAdd_Success(t *testing.T) {
	t.Run("Calculator Add operator testing", func(t *testing.T) {
		var num1 float64 = 10
		var num2 float64 = 20
		calc := Calculator{Num1: num1, Num2: num2}
		r, err := calc.Add()
		if err != nil {
			// fungsi built-in go dari package log
			// jika terjadi error maka akan exit
			log.Fatalln(err)
		}
		if *r != 30 {
			t.Error("It should return 30")
		}
	})
}

func TestCalculatorAdd_Fail(t *testing.T) {
	t.Run("Calculator Add detected negative number testing", func(t *testing.T) {
		var num1 float64 = -10
		var num2 float64 = 20
		calc := Calculator{Num1: num1, Num2: num2}
		_, err := calc.Add()
		if err != nil && err.Error() != "negative number detected" {
			t.Error("It have a negative number")
		}
	})
}

func TestCalculatorSub_Success(t *testing.T) {
	t.Run("Calculator Sub operator testing", func(t *testing.T) {
		var num1 float64 = 10
		var num2 float64 = 5
		calc := Calculator{Num1: num1, Num2: num2}
		r, err := calc.Sub()
		if err != nil {
			// fungsi built-in go dari package log
			// jika terjadi error maka akan exit
			log.Fatalln(err)
		}
		if *r != 5 {
			t.Error("It should return 5")
		}
	})
}

func TestCalculatorSub_Fail(t *testing.T) {
	t.Run("Calculator Sub detected negative number testing", func(t *testing.T) {
		var num1 float64 = -10
		var num2 float64 = 5
		calc := Calculator{Num1: num1, Num2: num2}
		_, err := calc.Sub()
		if err != nil && err.Error() != "negative number detected" {
			t.Error("It have a negative number")
		}
	})
}
