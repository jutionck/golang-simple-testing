package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CalculatorTestSuite struct {
	suite.Suite
	calService Calculator
}

// SetupTest -> otomatis dijalankan
func (suite *CalculatorTestSuite) SetupTest() {
	suite.calService = Calculator{}
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Success() {
	var num1 float64 = 10
	var num2 float64 = 20
	calc := Calculator{Num1: num1, Num2: num2}
	r, err := calc.Add()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), 30.0, *r)
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Fail() {
	var num1 float64 = -10
	var num2 float64 = 20
	calc := Calculator{Num1: num1, Num2: num2}
	_, err := calc.Add()
	require.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "negative number detected")
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Success() {
	var num1 float64 = 10
	var num2 float64 = 5
	calc := Calculator{Num1: num1, Num2: num2}
	r, err := calc.Sub()
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), 5.0, *r)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Fail() {
	var num1 float64 = -10
	var num2 float64 = 5
	calc := Calculator{Num1: num1, Num2: num2}
	_, err := calc.Sub()
	require.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "negative number detected")
}
