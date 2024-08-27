package main

import (
	"fmt"
	"sync"
)

func withSyncMap() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func withLockedMap() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	var lock sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			m[i] = i
			lock.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}

func withMap() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}

func main() {
	withSyncMap()
	withLockedMap()
}
