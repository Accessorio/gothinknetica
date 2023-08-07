package main

import (
	"flag"
	"fmt"
	"go-core-4/homework3/pkg/crawler"
	"go-core-4/homework3/pkg/crawler/spider"
	"go-core-4/homework3/pkg/index"
	"go-core-4/homework3/pkg/search"
)

func main() {
	f := flag.String("s", "", "word")
	flag.Parse()
	var id int
	var b []crawler.Document
	var array = [...]string{"https://go.dev", "https://golang.org"}
	for _, link := range array {
		s := spider.New()
		d, err := s.Scan(link, 2, &id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b = append(b, d...)

	}
	idb := index.Index(b)
	x := search.Search(idb, f)
	for _, id := range x {
		val := b[id]
		fmt.Println("id:", id, val.Title, "\n", val.URL)
	}
}
