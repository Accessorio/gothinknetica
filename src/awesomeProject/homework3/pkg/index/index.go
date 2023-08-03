package index

import (
	"go-core-4/homework3/pkg/crawler"
	"strings"
	"unicode"
)

type Index map[string][]int

func Indexing(a []crawler.Document) Index {
	idb := make(Index)
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
