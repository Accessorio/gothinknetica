package api

import (
	"encoding/json"
	"go-core-4/homework18/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Router  *mux.Router
	storage storage.Store
}

func New(storage storage.Store) *API {
	api := &API{
		Router:  mux.NewRouter(),
		storage: storage,
	}

	api.routes()

	return api
}

func (api *API) routes() {
	api.Router.HandleFunc("/add", api.addLinkHandler).Methods("POST")
	api.Router.HandleFunc("/get/{shortLink}", api.getLinkHandler).Methods("GET")
}

func (api *API) addLinkHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Link string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortLink := api.storage.Add(request.Link)
	response := map[string]string{"shortLink": shortLink}

	json.NewEncoder(w).Encode(response)
}

func (api *API) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortLink := vars["shortLink"]

	link, ok := api.storage.Get(shortLink)
	if !ok {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"link": link}
	json.NewEncoder(w).Encode(response)
}
