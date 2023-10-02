package main

import (
	"fmt"
)

func main() {
	fmt.Println(isdigitpermutation(2021, 1203))
}

func isdigitpermutation(x int, y int) bool {
	var countdigits = func(num int) [10]int {
		digitcount := [10]int{}
		for num > 0 {
			lastdigit := num % 10
			digitcount[lastdigit] += 1
			num /= 10
		}
		return digitcount
	}

	digitsx := countdigits(x)
	digitsy := countdigits(y)

	for digit := 0; digit < 10; digit++ {
		if digitsx[digit] != digitsy[digit] {
			return false
		}
	}
	return true
}
