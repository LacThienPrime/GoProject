package singleton

import (
	"reflect"
	"sync"
)

var (
	instances = make(map[string]any)
	lock      sync.RWMutex
)

func GetInstance[T any](constructor func() T) T {
	var v T
	name := reflect.TypeOf(v).String()

	v, available := func() (T, bool) {
		lock.RLock()
		defer lock.RUnlock()

		v, available := instances[name].(T)
		return v, available
	}()

	if available {
		return v
	}

	lock.Lock()
	defer lock.Unlock()

	v, available = instances[name].(T)
	if available {
		return v
	}

	v = constructor()
	instances[name] = v

	return v
}
