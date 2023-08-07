package main

import (
	"flag"
	"fmt"
	"go-core-4/homework5/pkg/crawler"
	"go-core-4/homework5/pkg/crawler/spider"
	"go-core-4/homework5/pkg/file"
	"go-core-4/homework5/pkg/index"
	"go-core-4/homework5/pkg/search"
	"log"
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
				continue
			}
			b = append(b, d...)

		}
		cdoc, err := file.CreateFile("./crawlerDoc.json")
		if err != nil {
			log.Fatal("Failed to Create:", err)
		}
		err = file.WriteToFile(b, cdoc)
		if err != nil {
			log.Fatal("Failed to Write to File:", err)
		}

		idb := index.Index(b)
		idoc, err := file.CreateFile("./indexDoc.json")
		if err != nil {
			log.Fatal("Failed to Create:", err)
		}
		err = file.WriteToFile(idb, idoc)

	default:
		idbf, err := file.OpenFile("indexDoc.json")
		if err != nil {
			log.Fatal("Failed to Open:", err)
		}
		idb, err := file.ReadFromIndex(idbf)
		if err != nil {
			log.Fatal("Failed to Read:", err)
		}
		x := search.Searching(idb, f)
		crawf, err := file.OpenFile("crawlerDoc.json")
		if err != nil {
			log.Fatal("Failed to Open:", err)
		}
		b, err := file.ReadFromCrawler(crawf)
		if err != nil {
			log.Fatal("Failed to Read:", err)
		}
		for _, id := range x {
			val := b[id]
			fmt.Println("id:", id, val.Title, "\n", val.URL)
		}

	}

}
