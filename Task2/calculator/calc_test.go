package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTidyString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(TidyString("wqweas 23 zxcz + 23"), "23+23")

	assert.Equal(TidyString("kek (23 zxcz + 23) * lol 5"), "(23+23)*5")
}
func TestCheckBasics(t *testing.T) {
	assert := assert.New(t)

	res, err := Calculator([]string{"1", "+", "2"})
	if assert.NoError(err) {
		assert.Equal(res, "3")
	}

	res, err = Calculator([]string{"8", "-", "4"})
	if assert.NoError(err) {
		assert.Equal("4", res)
	}

	res, err = Calculator([]string{"6", "/", "3"})
	if assert.NoError(err) {
		assert.Equal("2", res)
	}

	res, err = Calculator([]string{"100", "*", "5"})
	if assert.NoError(err) {
		assert.Equal("500", res)
	}
}

func TestCheckFail(t *testing.T) {
	assert := assert.New(t)

	res, err := Calculator([]string{"1", "+", "-", "2"})
	if assert.NoError(err) {
		assert.NotEqual("3", res)
	}

	res, err = Calculator([]string{"8", "-", "4", "5", "2"})
	if assert.NoError(err) {
		assert.NotEqual("4", res)
	}
}

func TestParenthesis(t *testing.T) {
	assert := assert.New(t)

	res, err := Calculator([]string{"(", "1", "+", "2", ")", "-", "3"})
	if assert.NoError(err) {
		assert.Equal("0", res)
	}

	res, err = Calculator([]string{"(", "1", "+", "2", ")", "*", "3"})
	if assert.NoError(err) {
		assert.Equal("9", res)
	}

	res, err = Calculator([]string{"(", "1", "2", "-", "5", ")", "*", "(", "1", "2", "+", "3", ")"})
	if assert.NoError(err) {
		assert.Equal("105", res)
	}
}
