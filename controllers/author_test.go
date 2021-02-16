package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNewAuthor(t *testing.T) {
	newAuthor := NewAuthor()
	assert.NotNil(t, newAuthor)
}