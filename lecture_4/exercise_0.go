package main

import (
	"fmt"
	"slices"
)

func main() {
	sec := []int{5, 5, 3, 4, 4, 5}
	fmt.Println(countsort(sec))
}

func countsort(sec []int) []int {
	minval := slices.Min(sec)
	maxval := slices.Max(sec)
	k := (maxval - minval + 1)
	count := make([]int, k)
	for _, now := range sec {
		count[now-minval] += 1
	}
	nowpos := 0
	for val := 0; val < k; val++ {
		for i := 0; i < count[val]; i++ {
			sec[nowpos] = val + minval
			nowpos += 1
		}
	}
	return sec
}
