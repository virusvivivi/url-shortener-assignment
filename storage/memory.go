// storage: lưu tạm bộ nhớ
package storage

import (
	"sync"

	"urlshortener/model"
)

type Store interface {
	Save(link *model.Shortlink) error
	FindByID(id string) *model.Shortlink
	FindByOriginalURL(url string) *model.Shortlink
	Close() error
}

type memoryStore struct {
	mu    sync.RWMutex
	byID  map[string]*model.Shortlink
	byURL map[string]*model.Shortlink
}

func NewMemoryStore() Store {
	return &memoryStore{
		byID:  make(map[string]*model.Shortlink),
		byURL: make(map[string]*model.Shortlink),
	}
}

func (s *memoryStore) Save(link *model.Shortlink) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.byID[link.ID] = link
	s.byURL[link.OriginalURL] = link
	return nil
}

func (s *memoryStore) FindByID(id string) *model.Shortlink {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.byID[id]
}

func (s *memoryStore) FindByOriginalURL(url string) *model.Shortlink {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.byURL[url]
}

func (s *memoryStore) Close() error {
	return nil
} 