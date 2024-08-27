### Slice 的容量与重分配

在使用 slices 时，如果超出了它的容量，Go 会自动重新分配内存，这可能会导致性能问题。

```go
package main

import "fmt"

func main() {
    var s []int
    for i := 0; i < 1000000; i++ {
        s = append(s, i)
    }
    fmt.Println("Length:", len(s))
    fmt.Println("Capacity:", cap(s)) // 容量可能会比预期的大很多
}
```

#### 修正方法：
预先分配足够的容量以减少内存重分配的次数。

```go
package main

import "fmt"

func main() {
    s := make([]int, 0, 1000000) // 预先分配容量
    for i := 0; i < 1000000; i++ {
        s = append(s, i)
    }
    fmt.Println("Length:", len(s))
    fmt.Println("Capacity:", cap(s)) // 容量应与预期一致
}
```

#### 验证代码：

```go
func fromEmptySlice() []int {
	var s []int
	for i := 0; i < 1000000; i++ {
		s = append(s, i)
	}
	return s
}

func fromPreAllocatedSlice() []int {
	s := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		s = append(s, i)
	}
	return s
}
```

#### 分配空间对比：
```bash
From empty, length: 1000000
From empty, capacity: 1055744
From pre-allocated, length: 1000000
From pre-allocated, capacity: 1000000
```

可以看到，从空 slice 开始，最终的容量比预期的大很多，而从预先分配的 slice 开始，容量与预期一致。

#### 性能对比：

```bash
BenchmarkFromEmptySlice-16           	    4413	   2612088 ns/op
BenchmarkFromPreAllocatedSlice-16    	   28468	    416781 ns/op
```

从测试结果可以看出，预先分配版本性能提升了 6 倍。