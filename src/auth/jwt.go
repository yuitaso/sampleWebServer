package auth

import (
	"crypto/rsa"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/env"
	"github.com/yuitaso/sampleWebServer/src/util"
)

func getPrivateKey() *rsa.PrivateKey { // TODO 初回のみ読み込みに変更
	path, err := filepath.Abs(env.PrivateKeyPath)
	util.MustNot(err)
	bytes, err := os.ReadFile(path)
	util.MustNot(err)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	util.MustNot(err)

	return key
}

func getPublicKey() *rsa.PublicKey { // TODO 初回のみ読み込みに変更
	path, err := filepath.Abs(env.PublicKeyPath)
	util.MustNot(err)
	bytes, err := os.ReadFile(path)
	util.MustNot(err)
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	util.MustNot(err)

	return key
}

type AuthClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User) (string, error) {
	claims := AuthClaims{
		user.Uuid.String(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "me", // fix me
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key := getPrivateKey()
	tokenstring, err := token.SignedString(key)
	fmt.Println(tokenstring)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ParseToken(requestToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		requestToken,
		func(t *jwt.Token) (interface{}, error) {
			return getPublicKey(), nil
		})
	if err != nil {
		return &jwt.Token{}, err
	}

	return token, nil
}
