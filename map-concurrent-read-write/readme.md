### Map 并发读写问题

在多个 goroutine 中并发读写 map 会导致崩溃。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
    fmt.Println(m) // 可能会崩溃
}
```

多执行几次，就会看到崩溃的问题：

```shell
❯ go run main.go
fatal error: concurrent map writes

goroutine 27 [running]:
main.main.func1(0x9)
	/Users/kevin/Develop/go/opensource/kevwan/go-best-practices/map-concurrent-read-write/main.go:16 +0x60
created by main.main in goroutine 1
	/Users/kevin/Develop/go/opensource/kevwan/go-best-practices/map-concurrent-read-write/main.go:14 +0x48

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x140000021c0?)
	/Users/kevin/.gvm/gos/go1.22.5/src/runtime/sema.go:62 +0x2c
sync.(*WaitGroup).Wait(0x14000110020)
	/Users/kevin/.gvm/gos/go1.22.5/src/sync/waitgroup.go:116 +0x74
main.main()
	/Users/kevin/Develop/go/opensource/kevwan/go-best-practices/map-concurrent-read-write/main.go:20 +0xf8
exit status 2
```

#### 修正方法：

使用 `sync.Map`：

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
```

或者添加读写锁来保护 map：

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
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
```

#### 总结

在使用数据结构时，一定要考虑并发读写的问题，同时也要考虑锁粒度问题，特别避免锁 IO 的操作。