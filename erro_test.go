package erro

import (
	"fmt"
	"testing"
)

func TestErro(t *testing.T) {
	TracePkg("github.com/lincaiyong/erro")
	defer Recover(func(err error) {
		fmt.Println(err)
	})
	foo()
}

func foo() {
	Assert(false, "test")
}
