package main

import (
	"fmt"
	gf "github.com/nasan016/gofer"
)

func main() {
	x := gf.Ref("Hello")

	gf.WatchEffect(func() {
		fmt.Println(x.GetValue())
	})

	x.SetValue("World!")
}
