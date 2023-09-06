package main

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
	err := createDocs()
	if err != nil {
		log.Fatal(err)
	}
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
		m, err := find(&s)
		if err != nil {
			log.Fatal(err)
		}
		err = netsrv.ServerRequest(conn, m)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func createDocs() error {
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
		log.Printf("Failed to Create: %v", err)
		return err
	}
	err = file.WriteToFile(b, cdoc)
	if err != nil {
		log.Printf("Failed to Write to File: %v", err)
		return err
	}

	idb := index.Index(b)
	idoc, err := file.CreateFile("./indexDoc.json")
	if err != nil {
		log.Printf("Failed to Create: %v", err)
		return err
	}
	err = file.WriteToFile(idb, idoc)
	if err != nil {
		log.Printf("Failed to Write to file: %v", err)
		return err
	}
	return nil
}

func find(f *string) ([]byte, error) {
	idbf, err := file.OpenFile("indexDoc.json")
	if err != nil {
		log.Printf("Failed to Open: %v", err)
		return nil, err
	}
	idb, err := file.ReadFromIndex(idbf)
	if err != nil {
		log.Printf("Failed to Read: %v", err)
		return nil, err
	}
	x := search.Searching(idb, f)
	crawf, err := file.OpenFile("crawlerDoc.json")
	if err != nil {
		log.Printf("Failed to Open: %v", err)
		return nil, err
	}
	b, err := file.ReadFromCrawler(crawf)
	if err != nil {
		log.Printf("Failed to Read: %v", err)
		return nil, err
	}
	m := make([]byte, 1)
	for _, id := range x {
		val := b[id]
		a := []byte(val.Title + ", " + val.URL + "\n")
		m = append(m, a...)
	}
	m = append(m, byte('*'))
	return m, nil
}
