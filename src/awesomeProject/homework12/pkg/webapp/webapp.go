package webapp

import (
	"encoding/json"
	"go-core-4/homework12/pkg/crawler"
	"go-core-4/homework12/pkg/file"
	"net/http"
	"sync"
)

type IndexerData struct {
	Word    string `json:"token"`
	Indexes []int  `json:"list"`
}

type CrawlerData struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}

type API struct {
	crawler []crawler.Document
	indexer map[string][]int
	rwm     sync.RWMutex
}

func Fill(c []crawler.Document, m map[string][]int) *API {
	return &API{
		crawler: c,
		indexer: m,
	}
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
	idbf, err := file.OpenFile("indexDoc.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	idb, err := file.ReadFromIndexJSON(idbf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rslt, err := json.Marshal(idb)

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
	crawf, err := file.OpenFile("crawlerDoc.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cra, err := file.ReadFromCrawlerJSON(crawf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rslt, err := json.Marshal(cra)
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
