package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func goRoutines() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.RUnlock()
		m.Lock()
		go incrementer()
		m.Unlock()
	}
	wg.Wait()
}

func incrementer() {
	counter++
	wg.Done()
}
func sayHello() {
	fmt.Printf("HELLO %d!\n", counter)
	wg.Done()
}
