package search

import (
	"fmt"
	Index "go-core-4/homework5/pkg/index"
	"strings"
)

func Searching(a Index.Index, s *string) []int {
	fnd := strings.ToLower(*s)
	var x []int
	for _, c := range Index.SToWords(fnd) {
		if b, ok := a[c]; ok {
			if x == nil {
				x = b
			} else {
				x = search(x, b)
			}
		} else {
			fmt.Println("Nothing was find.")
			break
		}
	}
	return x
}

func search(x []int, y []int) []int {
	maxLen := len(x)
	if len(y) > maxLen {
		maxLen = len(y)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			i++
		} else if x[i] > y[j] {
			j++
		} else {
			r = append(r, x[i])
			i++
			j++
		}
	}
	return r
}
