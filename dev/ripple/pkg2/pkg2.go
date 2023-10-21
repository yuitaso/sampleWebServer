package pkg2

import (
	"fmt"
	"github.com/yuitaso/sampleWebServer/dev/ripple/pkg"
)

func Foo() string {
	return fmt.Sprintf("Foo ", pkg.Hoge)
}
