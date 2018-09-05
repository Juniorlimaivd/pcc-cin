package main

import (
	"fmt"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func createBorderArray(pattern string) []int {
	m := len(pattern)
	var next [100]int
	i := 1
	j := 0

	// aaaabbbb
	// ababcabab

	for i < m {
		for i+j < m && pattern[i+j] == pattern[j] {
			next[i+j+1] = j + 1
			j++
		}

		next[i+j+1] = next[j]

		i += Max(1, j-next[j])
		j = Max(0, next[j])
	}

	return next[:m+1]

}

func kmp(pattern string, text string) []int {
	next := createBorderArray(pattern)
	fmt.Println(next)
	res := []int{}
	n := len(text)
	m := len(pattern)
	j := 0
	i := 0
	for i <= n-m {
		for j < m && text[i+j] == pattern[j] {
			j++
		}

		if j == m {
			res = append(res, i)
		}

		i += Max(1, j-next[j])
		j = Max(0, next[j])

	}

	return res
}

func main() {

	t := `Those hours that with gentle work did frame
	The lovely gaze where every eye doth dwell
	Will play the tyrants to the very same,
	And that unfair which fairly doth excel:
	For never-resting time leads summer on
	To hideous winter and confounds him there,
	Sap checked with frost and lusty leaves quite gone,
	Beauty o'er-snowed and bareness every where:
	Then were not summer's distillation left
	A liquid prisoner pent in walls of glass,
	Beauty's effect with beauty were bereft,
	Nor it nor no remembrance what it was.
	  But flowers distilled though they with winter meet,
	  Leese but their show, their substance still lives sweet.
  `
	p := "eye"

	r := kmp(p, t)
	fmt.Println(r)

}
