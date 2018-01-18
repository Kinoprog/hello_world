package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBla(t *testing.T) {
	x := Bla("s")
	assert.True(t, x == "b", "fue %s", x)
}
