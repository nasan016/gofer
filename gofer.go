package gofer

import (
	"reflect"
)

type effect func()

var currentActiveEffect effect

type dep struct {
	subscribers map[uintptr]effect
}

func newDep() *dep {
	return &dep{subscribers: make(map[uintptr]effect)}
}

func (d *dep) track(update effect) {
	key := reflect.ValueOf(update).Pointer()
	d.subscribers[key] = currentActiveEffect
}

func (d *dep) trigger() {
	for _, effect := range d.subscribers {
		if effect != nil {
			effect()
		}
	}
}

/*
refImpl (ref) is a reactive primitive that can be read and written onto
*/
type refImpl[T any] struct {
	dep   *dep
	value T
}

func Ref[T any](initialValue T) *refImpl[T] {
	return &refImpl[T]{
		dep:   newDep(),
		value: initialValue,
	}
}

func (r *refImpl[T]) GetValue() T {
	r.dep.track(currentActiveEffect)
	return r.value
}

func (r *refImpl[T]) SetValue(newValue T) {
	r.value = newValue
	r.dep.trigger()
}

/*
computed is a ref that is computed by a getter
*/
type computed[T any] struct {
	dep     *dep
	compute func() T
}

func Computed[T any](computedValue func() T) *computed[T] {

	return &computed[T]{
		dep:     newDep(),
		compute: computedValue,
	}
}

func (c *computed[T]) GetValue() T {
	c.dep.track(currentActiveEffect)
	return c.compute()
}

/*
Runs a function immediately while reactively tracking its dependencies
and re-runs it whenever the dependencies are changed.
*/
func WatchEffect(update effect) {
	var wrappedUpdate func()
	wrappedUpdate = func() {
		currentActiveEffect = wrappedUpdate
		update()
	}
	wrappedUpdate()
}
