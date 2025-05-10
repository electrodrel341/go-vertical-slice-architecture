package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
)

func main2() {
	privateKeyPath := "config/jwt_private.key" // путь к приватному ключу
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("failed to read private key: %v", err))
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %v", err))
	}

	// --- Claims (можешь менять под себя) ---
	claims := jwt.MapClaims{
		"user_id": "e1a0633e-04c6-4d92-8a12-10f3f1df0000",
		"role":    "admin", // или "user", "manager"
		"exp":     time.Now().Add(1500 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signed, err := token.SignedString(privateKey)
	if err != nil {
		panic(fmt.Sprintf("failed to sign token: %v", err))
	}

	fmt.Println("Generated JWT:\n")
	fmt.Println(signed)
}
