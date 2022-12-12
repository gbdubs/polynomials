package precision

import "sync"

const defaultPrecision uint = 1000

var precision uint = defaultPrecision
var lock sync.RWMutex

func Set(p uint) {
	lock.Lock()
	precision = p
	lock.Unlock()
}

func Get() uint {
	return precision
}
