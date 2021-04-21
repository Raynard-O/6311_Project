package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Secrets struct {
	MONGO_ENV			string
	MONGO_HOST         string
	MONGO_DB           string
	MONGO_USER         string
	MONGO_PASS         string
	MONGO_PORT         string
	Port               string
	Environment        string
	DatabaseURL        string
	PulsarUrl          string
	DatabaseName       string
	ServiceName        string
	ComplyAdvantageURL string
	ComplyAdvantageAPI string
	HmacSigningKey     string
	mu                 *sync.Mutex
	GoogleCientID      string
	GoogleSecret       string
}





const (
	ServiceName = "comply"
	Domain      = "io.roava"
	Local       = "local"
	Production  = "production"
)

var EventRoot = fmt.Sprintf("%s.%s", Domain, ServiceName)

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}



// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets( debug bool) (*Secrets, error) {


	//fmt.Print(debug)
	if !debug {
		err := Load()
		if err != nil{
			return nil, err
		}
	}

	_secrets := &Secrets{
		MONGO_ENV:         	os.Getenv("GO_ENV"),
		mu:                 &sync.Mutex{},
		MONGO_HOST:         os.Getenv("MONGO_HOST"),
		MONGO_PORT:         os.Getenv("MONGO_PORT"),
		MONGO_DB:           os.Getenv("MONGO_DB"),
		MONGO_USER:         os.Getenv("MONGO_USER"),
		MONGO_PASS:         os.Getenv("MONGO_PASS"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		DatabaseName:       os.Getenv("DATABASE_NAME"),
		Environment:        os.Getenv("ENVIRONMENT"),
		PulsarUrl:          os.Getenv("PULSAR_URL"),
		ServiceName:        os.Getenv("SERVICE_NAME"),
		ComplyAdvantageURL: os.Getenv("COMPLY_ADVANTAGE_URL"),
		ComplyAdvantageAPI: os.Getenv("COMPLY_ADVANTAGE_APIKEY"),
		HmacSigningKey:     os.Getenv("Hmac_Signing_Key"),
		GoogleCientID:      os.Getenv("Google_Cient_ID"),
		GoogleSecret:       os.Getenv("Google_Secret"),
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_secrets.Port = port
	return _secrets, nil
}

//type debugContext string
//
//func distance(ctx context.Context, k debugContext) *bool {
//	v := ctx.Value(k).(*bool)
//		return v
//}