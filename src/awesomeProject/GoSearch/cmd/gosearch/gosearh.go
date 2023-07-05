package main

import (
	"flag"
	"fmt"
	"go-core-4/GoSearch/pkg/crawler"
	"go-core-4/GoSearch/pkg/crawler/spider"
	"strings"
)

func main() {
	f := flag.String("s", "", "word")
	flag.Parse()
	var b []crawler.Document
	var array = [...]string{"https://go.dev", "https://golang.org"}
	for _, link := range array {
		s := spider.New()
		d, err := s.Scan(link, 2)
		if err != nil {
			fmt.Println(err)
		}
		b = append(b, d...)
		for _, c := range d {
			if *f != "" {
				if strings.Contains(c.Title, *f) {
					fmt.Println(c.URL)
				}

			}
		}
	}
}
