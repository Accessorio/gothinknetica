package file

import (
	"encoding/json"
	"fmt"
	"go-core-4/homework12/pkg/crawler"
	"io"
	"os"
)

func CreateFile(s string) (*os.File, error) {
	doc, err := os.Create(s)
	if err != nil {
		fmt.Println("Error in CreateFile", err)
		return nil, err
	}
	return doc, err
}

func WriteToFile(b any, f io.WriteCloser) error {
	defer f.Close()
	bytes, err := toBytes(b)
	if err != nil {
		fmt.Println("Error in writing file", err)
		return err
	}
	_, err = f.Write(bytes)
	if err != nil {
		fmt.Println("Error in writing file", err)
		return err
	}
	return err
}

func ReadFromCrawler(f io.ReadCloser) ([]crawler.Document, error) {
	defer f.Close()
	var craw []crawler.Document
	r, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error in reading crawler", err)
		return nil, err
	}
	err = json.Unmarshal(r, &craw)
	return craw, err
}

func ReadFromIndex(f io.ReadCloser) (map[string][]int, error) {
	defer f.Close()
	var ind map[string][]int
	r, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error in reading crawler", err)
		return nil, err
	}
	err = json.Unmarshal(r, &ind)
	return ind, err
}

func OpenFile(s string) (*os.File, error) {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("File doesn't exist", err)
		return nil, err
	}
	return f, err
}

func toBytes(b any) ([]byte, error) {
	bytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error in crawler.Document to []byte", err)
		return nil, err
	}
	return bytes, err
}
