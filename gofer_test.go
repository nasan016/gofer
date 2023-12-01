package gofer

import "testing"

// test to see if WatchEffect() re-evaluates expression
func TestWatchEffect(t *testing.T) {
	refBool := Ref(false)
	x := refBool.GetValue()

	WatchEffect(func() {
		x = refBool.GetValue()
	})

	refBool.SetValue(true)

	if !x {
		t.Fatalf(`x = %t, want true`, x)
	}
}

func TestComputed(t *testing.T) {
	x := Ref(1)
	y := Ref(1)
	z := Computed(func() int {
		return x.GetValue() + y.GetValue()
	})

	if z.GetValue() != x.GetValue()+y.GetValue() {
		t.Fatalf("err")
	}

	x.SetValue(2)

	if z.GetValue() != x.GetValue()+y.GetValue() {
		t.Fatalf("err")
	}
}
