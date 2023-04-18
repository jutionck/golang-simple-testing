package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculatorAdd_Success(t *testing.T) {
	t.Run("Calculator Add operator testing", func(t *testing.T) {
		var num1 float64 = 10
		var num2 float64 = 20
		calc := Calculator{Num1: num1, Num2: num2}
		r, err := calc.Add()
		require.NoError(t, err)
		assert.Equal(t, 30.0, *r)
	})
}

func TestCalculatorAdd_Fail(t *testing.T) {
	t.Run("Calculator Add detected negative number testing", func(t *testing.T) {
		var num1 float64 = -10
		var num2 float64 = 20
		calc := Calculator{Num1: num1, Num2: num2}
		_, err := calc.Add()
		require.Error(t, err)
		assert.EqualError(t, err, "negative number detected")
	})
}

func TestCalculatorSub_Success(t *testing.T) {
	t.Run("Calculator Sub operator testing", func(t *testing.T) {
		var num1 float64 = 10
		var num2 float64 = 5
		calc := Calculator{Num1: num1, Num2: num2}
		r, err := calc.Sub()
		require.NoError(t, err)
		assert.Equal(t, 5.0, *r)
	})
}

func TestCalculatorSub_Fail(t *testing.T) {
	t.Run("Calculator Sub detected negative number testing", func(t *testing.T) {
		var num1 float64 = -10
		var num2 float64 = 5
		calc := Calculator{Num1: num1, Num2: num2}
		_, err := calc.Sub()
		require.Error(t, err)
		assert.EqualError(t, err, "negative number detected")
	})
}
