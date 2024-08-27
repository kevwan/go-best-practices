package main

import "fmt"

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

func main() {
	s1 := fromEmptySlice()
	s2 := fromPreAllocatedSlice()
	fmt.Println("From empty, length:", len(s1))
	fmt.Println("From empty, capacity:", cap(s1))
	fmt.Println("From pre-allocated, length:", len(s2))
	fmt.Println("From pre-allocated, capacity:", cap(s2))
}
