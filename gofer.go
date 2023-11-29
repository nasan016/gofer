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
		effect()
	}
}

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

type computed[T any] struct {
}

func WatchEffect(update effect) {
	var wrappedUpdate func()
	wrappedUpdate = func() {
		currentActiveEffect = wrappedUpdate
		update()
	}
	wrappedUpdate()
}
