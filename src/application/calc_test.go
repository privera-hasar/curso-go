package calc_test

import (
	. "curso/application"
	"testing"

	"gotest.tools/assert"
)

func TestCalcAdd(t *testing.T) {
	t.Parallel()

	expected := 3

	sut := &Calculator{}

	actual := sut.Add(1, 2)

	assert.Equal(t, expected, actual)
}

func TestCalcDivide(t *testing.T) {
	t.Parallel()

	expected := 2

	sut := &Calculator{}

	actual, err := sut.Divide(6, 3)

	assert.Assert(t, err == nil, err)
	assert.Equal(t, expected, actual)
}

func TestCalcDivideByZero(t *testing.T) {
	t.Parallel()

	sut := &Calculator{}

	_, err := sut.Divide(6, 0)

	assert.Assert(t, err != nil)
}
