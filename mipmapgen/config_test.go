package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseConfigure(t *testing.T) {
	configure, err := parseConfigure("examples/")
	assert.NoError(t, err)
	assert.NotNil(t, configure)
	assert.NotNil(t, configure.Requests)
	assert.NotEmpty(t, configure.Requests)
}
