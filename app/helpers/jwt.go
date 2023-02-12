package helpers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID            string    `json:"id"`
	Nis           string    `json:"nis"`
	Nisn          string    `json:"nisn"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	Status        bool      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Exp           int64     `json:"exp"`
	Role          string    `json:"role"`
	ClientID      string    `json:"client_id"`
	Authorization bool      `json:"authorization"`
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecrePublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {

	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}
	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(LoadEnv(SecrePublicKeyEnvName)))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecrePublicKeyEnvName, requestToken string) (*AccessToken, error) {

	jwtSecretKey := LoadEnv(SecrePublicKeyEnvName)
	token, err := jwt.Parse((requestToken), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	claimToken := DecodeToken(token)
	return &claimToken, nil
}

func VerifyToken(accessToken, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := LoadEnv(SecrePublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}
