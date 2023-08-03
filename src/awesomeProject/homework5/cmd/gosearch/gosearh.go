package main

import (
	"flag"
	"fmt"
	"go-core-4/homework5/pkg/crawler"
	"go-core-4/homework5/pkg/crawler/spider"
	"go-core-4/homework5/pkg/file"
	"go-core-4/homework5/pkg/index"
	"go-core-4/homework5/pkg/search"
)

func main() {
	f := flag.String("s", "", "word")
	flag.Parse()
	switch name := *f; name {
	case "":
		var id int
		var b []crawler.Document
		var array = [...]string{"https://go.dev", "https://golang.org"}
		for _, link := range array {
			s := spider.New()
			d, err := s.Scan(link, 2, &id)
			if err != nil {
				fmt.Println(err)
			}
			b = append(b, d...)

		}
		cdoc := file.CreateFile("./crawlerDoc.json")
		file.WriteToFile(b, cdoc)
		idb := index.Indexing(b)
		idoc := file.CreateFile("./indexDoc.json")
		file.WriteToFile(idb, idoc)

	default:
		idbf := file.OpenFile("indexDoc.json")
		idb := file.ReadFromIndex(idbf)
		x := search.Searching(idb, f)
		crawf := file.OpenFile("crawlerDoc.json")
		b := file.ReadFromCrawler(crawf)
		for _, id := range x {
			val := b[id]
			fmt.Println("id:", id, val.Title, "\n", val.URL)
		}

	}

}
