// O(N) complexity

package main

import (
	"fmt"
	"slices"
)

func main() {
	x := 6
	myset := []int{3, 7, 2, 8, 4}
	fmt.Println(sum(x, myset...))
}

func sum(x int, myset ...int) (int, int) {
	prevnums := []int{}
	for _, nownum := range myset {
		if slices.Contains(prevnums, x-nownum) {
			return nownum, x - nownum
		}
		prevnums = append(prevnums, nownum)
	}
	return 0, 0
}
