package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReader(t *testing.T) {
	newReader := NewReader()
	assert.NotNil(t, newReader)
}
