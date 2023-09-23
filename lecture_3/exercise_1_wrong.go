package main

import "fmt"

func main() {
	x := 6
	myset := []int{1, 7, 4, 8, 2}
	fmt.Println(sum(x, myset...))
}

func sum(x int, myset ...int) (int, int) {
	for _, a := range myset {
		for _, b := range myset {
			if a+b == x {
				return a, b
			}
		}
	}
	return 0, 0
}
