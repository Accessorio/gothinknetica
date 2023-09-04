package webapp

import (
	"encoding/json"
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
	crawlerData []*CrawlerData
	indexerData []*IndexerData
	rwm         sync.RWMutex
}

func (a *API) Fill(c []*CrawlerData, m []*IndexerData) {
	a.rwm.Lock()
	a.crawlerData = c
	a.indexerData = m
	a.rwm.Unlock()
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
