package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func fatal(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func main() {
	path, err := filepath.Abs("./dev/secrets/secret.pem")
	fatal(err)
	bytes, err := os.ReadFile(path)
	fatal(err)

	block, _ := pem.Decode(bytes)

	var key *rsa.PrivateKey
	if block.Type == "RSA PRIVATE KEY" {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		fatal(err)
	} else {
		fatal(errors.New("Unexpected key type"))
	}

	key.Precompute()
	if err := key.Validate(); err != nil {
		log.Fatal("validation error")
	}

	fmt.Println("キー: ", key)
}
