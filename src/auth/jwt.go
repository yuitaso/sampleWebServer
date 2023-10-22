package auth

import (
	"crypto/rsa"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yuitaso/sampleWebServer/src/env"
	"github.com/yuitaso/sampleWebServer/src/util"
)

func getPrivateKey() *rsa.PrivateKey {
	path, err := filepath.Abs(env.PrivateKeyPath)
	util.MustNot(err)
	bytes, err := os.ReadFile(path)
	util.MustNot(err)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	util.MustNot(err)

	return key
}

func getPublicKey() *rsa.PublicKey {
	path, err := filepath.Abs(env.PublicKeyPath)
	util.MustNot(err)
	bytes, err := os.ReadFile(path)
	util.MustNot(err)
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	util.MustNot(err)

	return key
}

func GenerateToken() (string, error) {
	key := getPrivateKey()
	token := jwt.New(jwt.SigningMethodRS256)

	tokenstring, err := token.SignedString(key)
	fmt.Println(tokenstring)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func VelifyToken(requestToken string) error {
	token, err := jwt.Parse(
		requestToken,
		func(t *jwt.Token) (interface{}, error) {
			return getPublicKey(), nil
		})

	fmt.Println("とくん", token)

	return err
}
