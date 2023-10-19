package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashed_cost4, err := bcrypt.GenerateFromPassword([]byte("pass"), 4)
	hashed_cost10, err := bcrypt.GenerateFromPassword([]byte("pass"), 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("hashed4: ", hashed_cost4)
	fmt.Println("")
	fmt.Println("hashed4: ", string(hashed_cost4))
	fmt.Println("hashed10: ", hashed_cost10)
	fmt.Println("ここまでOK")

	err = bcrypt.CompareHashAndPassword(hashed_cost10, []byte("pass"))
	if err != nil {
		fmt.Println("エラー起きた")
		fmt.Println(err.Error())
	} else {
		//err = bcrypt.CompareHashAndPassword(hashed_cost10, []byte("pass"))
		fmt.Println("エラー起きてない")
	}
}
