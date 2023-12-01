package main

import (
	"fmt"
	gf "github.com/nasan016/gofer"
)

func main() {
	price := gf.Ref(2)
	quantity := gf.Ref(1000)

	revenue := gf.Computed(func() int {
		return price.GetValue() * quantity.GetValue()
	})

	gf.WatchEffect(func() {
		fmt.Println("revenue:", revenue.GetValue())
	})

	price.SetValue(price.GetValue() / 2)
	price.SetValue(price.GetValue() * 10)
	quantity.SetValue(quantity.GetValue() + 500)
}
