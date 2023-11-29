package gofer

type effect func()

var currentActiveEffect effect

type dep struct {
	subscribers []effect
}

func newDep() *dep {
	return &dep{}
}

func (d *dep) track() {
	d.subscribers = append(d.subscribers, currentActiveEffect)
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
	r.dep.track()
	return r.value
}

func (r *refImpl[T]) SetValue(newValue T) {
	r.value = newValue
	r.dep.trigger()
}

func WatchEffect(update effect) {
	var wrappedUpdate func()
	wrappedUpdate = func() {
		currentActiveEffect = wrappedUpdate
		update()
		currentActiveEffect = nil
	}
	wrappedUpdate()
}
