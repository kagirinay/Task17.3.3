package main

import (
	"fmt"
	"sync"
)

const step int64 = 1
const interationAmount int64 = 1000

func main() {
	var counter int64 = 0
	var m = sync.RWMutex{}
	var c = sync.NewCond(&m)
	checker := func() {
		defer m.RUnlock()
		for {
			m.RLock()
			if counter == interationAmount {
				c.Signal()
				break
			}
			m.RUnlock()
		}
	}
	increment := func() {
		c.L.Lock()
		counter += step
		c.L.Unlock()
	}
	for i := int64(1); i <= interationAmount; i++ {
		go increment()
	}
	go checker()
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
	fmt.Println(counter)
}
