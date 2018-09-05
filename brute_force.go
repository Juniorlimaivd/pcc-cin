package main

import (
	"fmt"
)

func bruteForceSearch(pattern string, text string) []int {
	m := len(pattern)
	n := len(text)
	ans := []int{}

	for i := 0; i <= n-m; i++ {
		j := 0

		for j < m && pattern[j] == text[i+j] {
			j = j + 1
		}

		if j == m {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	pat := "abra"
	text := "abracadabra"

	r := bruteForceSearch(pat, text)

	fmt.Print(r, "\n")
}
