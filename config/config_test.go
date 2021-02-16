package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadSecrets(t *testing.T) {

	secret, err := LoadSecrets()
	if err != nil{
		t.Error("error loading secrets from env")
	}
	assert.NotNil(t, secret.Port)
	assert.NotNil(t, secret.MONGO_HOST)
	assert.NotNil(t, secret.MONGO_PASS)
	assert.NotNil(t, secret.MONGO_PORT)
	assert.NotNil(t, secret.MONGO_USER)
	assert.NotNil(t, secret.HmacSigningKey)
}
