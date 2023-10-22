package auth

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey *[]byte

func init() {
	fmt.Println("読み込み始めるよー")

	path, err := filepath.Abs("./dev/secrets/secret.pem")
	if err != nil {
		log.Fatal(err.Error())
	}

	dat, err := os.ReadFile(path) // pem 形式で読み込み
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("cannnot open secret file.")
	}
	privateKey = &dat
	*privateKey = []byte("sample")
}

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS512) //とりまからのJWT

	tokenstring, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func VelifyToken(requestToken string) error {
	_, err := jwt.Parse(
		requestToken,
		func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, errors.New("Faild to pars jwt.")
			}
			return "", nil
		})

	return err
}
