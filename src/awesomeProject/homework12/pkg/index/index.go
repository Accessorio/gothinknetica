package index

import (
	"go-core-4/homework12/pkg/crawler"
	"strings"
	"unicode"
)

func Index(a []crawler.Document) map[string][]int {
	idb := make(map[string][]int)
	for _, c := range a {
		w := SToWords(strings.ToLower(c.Title))
		for _, v := range w {
			val := idb[v]
			if val != nil && val[len(val)-1] == c.ID {
				continue
			}
			idb[v] = append(val, c.ID)
		}
	}
	return idb
}

func SToWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
