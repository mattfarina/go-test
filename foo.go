package test

import (
	"fmt"

	"github.com/Masterminds/semver"
)

func Foo() {
	foo, err := semver.NewVersion("1.2.3")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(foo)
}
