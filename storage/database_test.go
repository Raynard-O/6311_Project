package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMongoInit(t *testing.T) {

	_, err  := mongoInit("6311", "author")

	//assert.NotEmpty(t, mongoDb)

	assert.Error(t, err, "Error init database")
}
