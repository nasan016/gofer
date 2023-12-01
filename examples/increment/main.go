package main

import (
	"fmt"
	gf "github.com/nasan016/gofer"
)

func main() {
	x := gf.Ref(0)
	y := gf.Ref(0)

	gf.WatchEffect(func() {
		fmt.Println("x: ", x.GetValue())
		fmt.Println("y: ", y.GetValue())
		fmt.Println("")
	})

	increment(x)
	increment(y)
	increment(x)
}

func increment(ref *gf.RefImpl[int]) {
	ref.SetValue(ref.GetValue() + 1)
}
