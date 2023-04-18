package service

import (
	"fmt"
	"math"
)

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (c *Calculator) Add() (*float64, error) {
	if math.Signbit(c.Num1) || math.Signbit(c.Num2) {
		return nil, fmt.Errorf("negative number detected")
	}
	result := c.Num1 + c.Num2
	return &result, nil
}

func (c *Calculator) Sub() (*float64, error) {
	if math.Signbit(c.Num1) || math.Signbit(c.Num2) {
		return nil, fmt.Errorf("negative number detected")
	}
	result := c.Num1 - c.Num2
	return &result, nil
}
