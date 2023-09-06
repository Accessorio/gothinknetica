package search

import (
	"fmt"

	"go-core-4/homework5/pkg/index"
	"strings"
)

func Searching(a map[string][]int, s *string) []int {
	fnd := strings.ToLower(*s)
	var x []int
	for _, c := range index.SToWords(fnd) {
		if b, ok := a[c]; ok {
			if x == nil {
				x = b
)

func Search(a index.Indexer, f *string) []int {
	fnd := strings.ToLower(*f)
	fmt.Println(fnd)
	var x []int
	for _, c := range index.SToWords(fnd) {
		if *f != "" {
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
		} else {
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
