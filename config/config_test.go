package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func TestLoadSecrets(t *testing.T) {
	// Set your envars to whatever you need for testing
	os.Setenv("GO_ENV", "production")
	os.Setenv("MONGO_HOST", "production")
	os.Setenv("MONGO_DB", "production")
	os.Setenv("MONGO_USER", "production")
	os.Setenv("MONGO_PASS", "production")
	os.Setenv("MONGO_PORT", "8080")
	os.Setenv("Hmac_Signing_Key", "production")


	secret, err := LoadSecrets()
	if err != nil{
		t.Errorf("error loading secrets from env: got error %v", err)
	}else{
		t.Log("Environmental Variable is parsed")
	}

	assert.NotEmpty(t, secret.MONGO_HOST)
	assert.NotEmpty(t, secret.MONGO_PASS)
	assert.NotEmpty(t, secret.MONGO_PORT)
	assert.NotEmpty(t, secret.MONGO_USER)
	assert.NotEmpty(t, secret.HmacSigningKey)
	assert.NotEmpty(t, secret.MONGO_ENV)
	fmt.Println(secret.MONGO_PORT)
}

//func TestLoad(t *testing.T) {
//	err := Load()
//	if err != nil{
//		t.Errorf("error loading secrets from env: got error %v", err)
//	}else{
//		t.Log("Environmental Variable is parsed")
//	}
//}