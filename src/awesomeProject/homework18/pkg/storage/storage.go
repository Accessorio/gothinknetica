package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
)

type Store interface {
	Add(link string) string
	Get(short string) (string, bool)
}

type mapStore struct {
	mu    sync.Mutex
	links map[string]string
}

func New() Store {
	return &mapStore{
		links: make(map[string]string),
	}
}

func (s *mapStore) Add(link string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	hash := sha256.Sum256([]byte(link))
	short := hex.EncodeToString(hash[:3])

	counter := 0
	for _, exists := s.links[short]; exists && counter < 1000; {
		counter++
		newLink := link + strconv.Itoa(counter)
		hash = sha256.Sum256([]byte(newLink))
		short = hex.EncodeToString(hash[:3])
	}

	s.links[short] = link
	return short
}

func (s *mapStore) Get(short string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	link, ok := s.links[short]
	return link, ok
}
