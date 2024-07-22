package index

import (
	"go-core-4/homework13/pkg/crawler"
	"strings"
	"unicode"
)

func Index(a map[int]crawler.Document) map[string][]int {
	idb := make(map[string][]int)
	for id, c := range a {
		w := SToWords(strings.ToLower(c.Title))
		for _, v := range w {
			val := idb[v]
			if val != nil && val[len(val)-1] == id {
				continue
			}
			idb[v] = append(val, id)
		}
	}
	return idb
}

func SToWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
