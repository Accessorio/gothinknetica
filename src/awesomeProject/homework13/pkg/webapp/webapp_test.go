package webapp

import (
	"fmt"
	"go-core-4/homework13/pkg/crawler"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAPI_Index(t *testing.T) {
	expected := `[{"token":"Go","list":"[1 2 3]"},{"token":"Language","list":"[4 5 6]"}]`
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	w := httptest.NewRecorder()
	var a API
	a.indexer = map[string][]int{"Go": []int{1, 2, 3}, "Language": []int{4, 5, 6}}
	var ind []*IndexerData
	for key, list := range a.indexer {
		ind = append(ind, &IndexerData{Word: key, Indexes: fmt.Sprintf("%v", list)})
	}
	a.indexerData = ind
	a.Index(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}

func TestAPI_Docs(t *testing.T) {
	expected := `[{"id":"0","title":"FTP RU","body":"","url":"https://ftp.ru/"},{"id":"1","title":"golang org","body":"","url":"https://golang-org.com/"}]`
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	w := httptest.NewRecorder()
	var a API
	test := []crawler.Document{
		{ID: 0,
			URL:   "https://ftp.ru/",
			Title: "FTP RU",
			Body:  "",
		},
		{
			ID:    1,
			URL:   "https://golang-org.com/",
			Title: "golang org",
			Body:  "",
		},
	}
	a.crawler = test
	var cra []*CrawlerData
	for _, val := range a.crawler {
		cra = append(cra, &CrawlerData{Id: strconv.Itoa(val.ID), Title: val.Title, Body: val.Body, URL: val.URL})
	}
	a.crawlerData = cra
	a.Docs(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
