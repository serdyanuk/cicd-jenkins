package inmemory

import (
	"sync"
)

type SearchHistoryRepo struct {
	list []string
	mu   sync.RWMutex
}

func NewSearchHistoryRepo() *SearchHistoryRepo {
	return &SearchHistoryRepo{
		list: make([]string, 0),
	}
}

func (s *SearchHistoryRepo) Add(str string) {
	s.mu.Lock()
	s.list = append(s.list, str)
	s.mu.Unlock()
}

func (s *SearchHistoryRepo) List() []string {
	s.mu.RLock()
	var cp = make([]string, len(s.list))
	copy(cp, s.list)
	s.mu.RUnlock()
	return cp
}

func (s *SearchHistoryRepo) Clean() {
	s.mu.Lock()
	s.list = make([]string, 0)
	s.mu.Unlock()
}
