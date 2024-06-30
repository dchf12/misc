package main

import "fmt"

func main() {
	var sum int
	Apply[int]([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum)

	ns := []int{1, 2, 3, 4}
	ms := Filter[int](ns, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms)

	var ss []string = Map([]int{10, 20}, func(n int) string {
		return fmt.Sprintf("0x%x", n)
	})
	fmt.Println(ss)

	var t *Tuple[int, string] = New(10, "hoge")
	fmt.Println(t.X, t.Y)
}

type Tuple[T1, T2 any] struct {
	X T1
	Y T2
}

func New[T1, T2 any](t1 T1, t2 T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		t1,
		t2,
	}
}

func Map[T1, T2 any](ns []T1, f func(T1) T2) []T2 {
	ret := make([]T2, len(ns))
	for i, v := range ns {
		ret[i] = f(v)
	}
	return ret
}

func Apply[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}

func Filter[T any](ns []T, filterFunc func(T) bool) []T {
	var result []T
	for _, v := range ns {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return result
}
