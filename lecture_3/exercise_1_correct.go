// O(N^2) complexity

package main

import "fmt"

func main() {
	x := 6
	myset := []int{3, 7, 2, 8, 4}
	fmt.Println(sum(x, myset...))
}

func sum(x int, myset ...int) (int, int) {
	for a := 0; a < len(myset); a++ {
		for b := (a + 1); b < len(myset); b++ {
			if myset[a]+myset[b] == x {
				return myset[a], myset[b]
			}
		}
	}
	return 0, 0
}
