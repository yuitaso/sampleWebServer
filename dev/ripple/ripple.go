package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("./dev/ripple/pkg/pkg.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(path)

	l, _ := os.ReadFile(path)
	fmt.Println(l)
}
