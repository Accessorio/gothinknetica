package file

import (
	"encoding/json"
	"fmt"
	"go-core-4/homework5/pkg/crawler"
	Index "go-core-4/homework5/pkg/index"
	"io"
	"os"
)

func CreateFile(s string) *os.File {
	doc, err := os.Create(s)
	if err != nil {
		fmt.Println("Error in CreateFile", err)
		os.Exit(1)
	}
	return doc
}

func WriteToFile(b any, f io.WriteCloser) {
	defer f.Close()
	bytes := toBytes(b)
	_, err := f.Write(bytes)
	if err != nil {
		fmt.Println("Error in writing file", err)
		os.Exit(1)
	}
}

func ReadFromCrawler(f io.ReadCloser) []crawler.Document {
	defer f.Close()
	var craw []crawler.Document
	r, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error in reading crawler", err)
		os.Exit(1)
	}
	err = json.Unmarshal(r, &craw)
	return craw
}

func ReadFromIndex(f io.ReadCloser) Index.Index {
	defer f.Close()
	var ind Index.Index
	r, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error in reading crawler", err)
		os.Exit(1)
	}
	err = json.Unmarshal(r, &ind)
	return ind
}

func OpenFile(s string) *os.File {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("File doesn't exist", err)
		os.Exit(1)
	}
	return f
}

func toBytes(b any) []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error in crawler.Document to []byte", err)
		os.Exit(1)
	}
	return bytes
}
