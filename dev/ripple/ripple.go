package main

import (
	"fmt"
	"flag"
	"github.com/yuitaso/sampleWebServer/dev/ripple/pkg"
	"github.com/yuitaso/sampleWebServer/dev/ripple/pkg2"
)

func init() {
	fmt.Println("initが走る")
	flag.Parse()
}

func main() {
	fmt.Println("FOO: ", pkg2.Foo())
	fmt.Println("なんらかの処理があってもう一度")
	fmt.Println("FOO: ", pkg2.Foo())
	fmt.Println("その後")
	fmt.Println("HOGE: ", pkg.Hoge)
}
