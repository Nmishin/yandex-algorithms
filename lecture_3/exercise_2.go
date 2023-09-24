package main

import (
	"fmt"
	"slices"
)

func main() {
	text := []string{"abc", "abcd"}
	dictionary := []string{"abcd"}
	fmt.Println(wordsindict(text, dictionary))
}

func wordsindict(text []string, dictionary []string) []string {
	goodwords := dictionary
	for _, word := range dictionary {
		for delpos, _ := range word {
			goodwords = append(goodwords, word[:delpos]+word[delpos+1:])
		}
	}
	ans := []string{}
	for _, word := range text {
		if slices.Contains(goodwords, word) {
			ans = append(ans, word)
		}
	}
	return ans
}
