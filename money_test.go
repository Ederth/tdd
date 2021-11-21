package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, NewDollar(5).Equals(NewDollar(6)))
	assert.False(t, NewDollar(5).Equals(NewFranc(5)))
}

func TestMultiplication(t *testing.T) {
	d := NewDollar(5)
	assert.Equal(t, NewDollar(10), d.Times(2))
	assert.Equal(t, NewDollar(15), d.Times(3))
}
