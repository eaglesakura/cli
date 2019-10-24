package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 0, Atoi("0"))
	assert.Equal(t, -1, Atoi("-1"))
	assert.Equal(t, 123, Atoi("123"))
}
