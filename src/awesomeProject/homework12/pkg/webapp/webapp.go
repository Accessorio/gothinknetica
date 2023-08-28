package webapp

import (
	"encoding/json"
	"fmt"
	"go-core-4/homework12/pkg/crawler"
	"net/http"
	"strconv"
	"sync"
)

type IndexerData struct {
	Word    string `json:"token"`
	Indexes string `json:"list"`
}

type CrawlerData struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}

type API struct {
	crawler     []crawler.Document
	indexer     map[string][]int
	crawlerData []*CrawlerData
	indexerData []*IndexerData
	rwm         sync.RWMutex
}

func (a *API) Fill(c []crawler.Document, m map[string][]int) {
	a.indexer = m
	a.crawler = c
	var ind []*IndexerData
	for key, list := range a.indexer {
		ind = append(ind, &IndexerData{Word: key, Indexes: fmt.Sprintf("%v", list)})
	}
	a.indexerData = ind
	var cra []*CrawlerData
	for _, val := range a.crawler {
		cra = append(cra, &CrawlerData{Id: strconv.Itoa(val.ID), Title: val.Title, Body: val.Body, URL: val.URL})
	}
	a.crawlerData = cra
}

func (a *API) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not supported.", 405)
		return
	}
	w.Write([]byte(`<html><body><h2>SIMPLE GO FIND App</h2></body></html>`))
}

func (a *API) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not supported.", 405)
		return
	}
	a.rwm.RLock()
	defer a.rwm.RUnlock()
	rslt, err := json.Marshal(a.indexerData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(rslt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) Docs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not supported.", 405)
		return
	}
	a.rwm.RLock()
	defer a.rwm.RUnlock()
	rslt, err := json.Marshal(a.crawlerData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(rslt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
