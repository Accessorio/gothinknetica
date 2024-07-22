package webapp

import (
	"go-core-4/homework13/pkg/crawler"
	"go-core-4/homework13/pkg/index"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPI_Read(t *testing.T) {
	want := `[{"URL":"go.ru","Title":"THE Go"},{"URL":"go.com","Title":"THE Go second"},{"URL":"go.md","Title":"THE Go Third"}]`
	r := httptest.NewRequest(http.MethodGet, "/read?text=Go", nil)
	w := httptest.NewRecorder()
	var a API
	a.crawler = map[int]crawler.Document{
		0: {"ya.ru ", "yandex"},
		1: {"go.ru", "THE Go"},
		2: {"go.com", "THE Go second"},
		3: {"go.md", "THE Go Third"},
		4: {"vk.ru", "Vkontakte"},
	}
	a.indexer = index.Index(a.crawler)
	a.Fill(a.crawler, a.indexer)
	a.Read(w, r)
	res := w.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected %s but got %v", want, string(got))
	}
}

func TestAPI_Delete(t *testing.T) {
	want := `Ok`
	r := httptest.NewRequest(http.MethodDelete, "/delete?id=1", nil)
	w := httptest.NewRecorder()
	var a API
	a.crawler = map[int]crawler.Document{
		0: {"ya.ru ", "yandex"},
		1: {"go.ru", "THE Go"},
		2: {"go.com", "THE Go second"},
		3: {"go.md", "THE Go Third"},
		4: {"vk.ru", "Vkontakte"},
	}
	a.indexer = index.Index(a.crawler)
	a.Fill(a.crawler, a.indexer)
	a.Delete(w, r)
	res := w.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected %s but got %v", want, string(got))
	}
}

func TestAPI_Update(t *testing.T) {
	want := `OK`

	r := httptest.NewRequest(http.MethodPost, "/update?id=1", strings.NewReader(`{"url":"sber.cloud" "titile":"SBER CLOUD"}`))
	w := httptest.NewRecorder()
	var a API
	a.crawler = map[int]crawler.Document{
		0: {"ya.ru ", "yandex"},
		1: {"go.ru", "THE Go"},
		2: {"go.com", "THE Go second"},
		3: {"go.md", "THE Go Third"},
		4: {"vk.ru", "Vkontakte"},
	}
	a.indexer = index.Index(a.crawler)
	a.Fill(a.crawler, a.indexer)
	a.Update(w, r)
	res := w.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected %s but got %v", want, string(got))
	}
}

func TestAPI_Create(t *testing.T) {
	want := `OK`

	r := httptest.NewRequest(http.MethodPost, "/update?id=12", strings.NewReader(`{"url":"sber.cloud" "titile":"SBER CLOUD"}`))
	w := httptest.NewRecorder()
	var a API
	a.crawler = map[int]crawler.Document{
		0: {"ya.ru ", "yandex"},
		1: {"go.ru", "THE Go"},
		2: {"go.com", "THE Go second"},
		3: {"go.md", "THE Go Third"},
		4: {"vk.ru", "Vkontakte"},
	}
	a.indexer = index.Index(a.crawler)
	a.Fill(a.crawler, a.indexer)
	a.Create(w, r)
	res := w.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error! %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected %s but got %v", want, string(got))
	}
}
