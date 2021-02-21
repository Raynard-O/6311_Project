package reader

import (
	"6311_Project/models"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/labstack/echo"
)

//var configuration, err = config.LoadSecrets(false)
//
//var HmacSigningKey = []byte(configuration.HmacSigningKey)

//func GenerateToken(user *models.User) (string, error) {
//	now := time.Now()
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
//		"user_id":   user.ID,
//		"email":     user.Email,
//		"issued_at": now.Unix(),
//		"expire_at": now.Add(time.Hour * 72).Unix(),
//	})
//	return token.SignedString(HmacSigningKey)
//}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func LoginUser(c echo.Context, user *models.Reader) error {
	token := "GenerateToken(user)"
	//channel := ComputeHmac256(string(user.Model.ID), GetHmacKey)
	return LoginUserResponse(c, user, token, user.Email)
}

//func GetUser(c echo.Context) *models.Author {
//	DB, err := storage.MongoInit("ICDE", "readers", context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	user := new(models.User)
//	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
//	DB.FindOneUser("users", "email", claims["email"].(string), &user)
//	//log.Println(user)
//	return user
//}
