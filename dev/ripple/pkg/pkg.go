package pkg

import "fmt"

var Hoge string

func init() {
	fmt.Println("init pkg")
	Hoge = "yyy"
}
