package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomHash() (string, error) {
	str := "dankpassword"

	hasher := sha256.New()
	hasher.Write([]byte(str))
	random, err := generateRandomBytes(256)
	if err != nil {
		return "", err
	}
	hasher.Write(random)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func MustNot(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
