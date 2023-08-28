package main

// Сервер текущего времени в соответсвии с RFC 867.

import (
	"bufio"
	"fmt"
	"go-core-4/homework11/pkg/crawler"
	"go-core-4/homework11/pkg/crawler/spider"
	"go-core-4/homework11/pkg/file"
	"go-core-4/homework11/pkg/index"
	"go-core-4/homework11/pkg/netsrv"
	"go-core-4/homework11/pkg/search"
	"log"
	"net"
)

func main() {
	createDocs()
	l, err := net.Listen("tcp4", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(conn)
	for {
		s, err := netsrv.ClientResponce(r)
		if err != nil {
			log.Fatal(err)
		}
		m := find(&s)
		err = netsrv.ServerRequest(conn, m)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func createDocs() {
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
	if err != nil {
		log.Fatal("Failed to Write to file:", err)
	}
}

func find(f *string) []byte {
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
	m := make([]byte, 1)
	for _, id := range x {
		val := b[id]
		a := []byte(val.Title + ", " + val.URL + "\n")
		m = append(m, a...)
	}
	m = append(m, byte('*'))
	return m
}
