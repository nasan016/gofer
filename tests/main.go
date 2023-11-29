package main

import (
	"fmt"
	gf "github.com/nasan016/gofer"
)

func main() {
	x := gf.Ref(1)

	gf.WatchEffect(func() {
		fmt.Println(x.GetValue())
	})

	x.SetValue(2)
	x.SetValue(3)
	x.SetValue(1)
	x.SetValue(x.GetValue() + 5)
}
