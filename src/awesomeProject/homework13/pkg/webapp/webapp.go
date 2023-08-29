package webapp

import (
	"encoding/json"
	"fmt"
	"go-core-4/homework13/pkg/crawler"
	"go-core-4/homework13/pkg/search"
	"io"
	"log"
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
	URL   string `json:"url"`
}

type API struct {
	crawler     map[int]crawler.Document
	indexer     map[string][]int
	crawlerJSON []*CrawlerData
	indexerJSON []*IndexerData
	rwm         sync.RWMutex
}

func (a *API) Fill(c map[int]crawler.Document, m map[string][]int) {
	a.indexer = m
	a.crawler = c
	var ind []*IndexerData
	for key, list := range a.indexer {
		ind = append(ind, &IndexerData{Word: key, Indexes: fmt.Sprintf("%v", list)})
	}
	a.indexerJSON = ind
	var cra []*CrawlerData
	for id, val := range a.crawler {
		cra = append(cra, &CrawlerData{Id: strconv.Itoa(id), Title: val.Title, URL: val.URL})
	}
	a.crawlerJSON = cra

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

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method is not supported.", 405)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	a.rwm.RLock()
	_, ok := a.crawler[id]
	if !ok {
		http.Error(w, "No record with this ID", 404)
		return
	}
	a.rwm.RUnlock()
	a.rwm.Lock()
	defer a.rwm.Unlock()
	delete(a.crawler, id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not supported.", 405)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	_, ok := a.crawler[id]
	if ok {
		http.Error(w, "Record with this ID is exist", 409)
		return
	}
	rqst, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in Request Body", 404)
	}
	var doc crawler.Document
	err = json.Unmarshal(rqst, &doc)

	a.rwm.Lock()
	defer a.rwm.Unlock()

	a.crawler[id] = doc
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is not supported.", 405)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}
	a.rwm.Lock()
	defer a.rwm.Unlock()
	_, ok := a.crawler[id]
	if !ok {
		http.Error(w, "No record with this ID", 409)
		return
	}

	rqst, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in Request Body", 404)
	}
	var doc crawler.Document
	err = json.Unmarshal(rqst, &doc)

	a.crawler[id] = doc
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method is not supported.", 405)
		return
	}

	t := r.URL.Query().Get("text")
	if t == "" {
		http.Error(w, "No record's", 409)
		return
	} else if t == "all" {
		a.rwm.RLock()
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(a.crawlerJSON)
		if err != nil {
			http.Error(w, "Server Error", 400)
			log.Println("Error in Marshal", err)
			return
		}
		w.Write(b)
		w.WriteHeader(http.StatusOK)
		return
	}
	a.rwm.RLock()
	defer a.rwm.RUnlock()
	id := search.Search(a.indexer, t)
	var db []crawler.Document
	for _, i := range id {
		db = append(db, a.crawler[i])
	}
	bd, err := json.Marshal(db)
	if err != nil {
		http.Error(w, "Server Error", 400)
		log.Println("Error in Marshal", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bd)
	w.WriteHeader(http.StatusOK)
}
