package logic

import (
	"log"
	structs "rto/struct"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var contextKey = "contextMap"
var Secret string
var SecretKey []byte

// var contextMap Payload
var INTERNAL = "INTERNAL"

func GenerateToken(userData *structs.User) (*string, error) {
	// setSecret()
	SecretKey = []byte("0fb7570f-0073-43b7-9f98-2061bb9c2349")
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["UserName"] = userData.Username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// claims["IsActive"] = userData.IsActive
	claims["userId"] = userData.UserId
	// claims["provider"] = userData.Provider
	// claims["services"] = userData.ServiceName
	// claims["role"] = userData.RoleId
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return nil, err
	}
	return &tokenString, nil
}
