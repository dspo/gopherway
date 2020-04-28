package singleton

import (
	"sync"
)

var o sync.Once

func GetInstance2() Singleton {
	o.Do(func() {
		instance = new(singleton)
	})
	return instance
}
