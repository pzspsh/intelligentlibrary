package mapsutil

import (
	"sync"
	"sync/atomic"

	errorutil "utils/errors"
)

var (
	ErrReadOnly = errorutil.New("map is currently in read-only mode").WithTag("syncLockMap")
)

// SyncLock adds sync and lock capabilities to generic map
type SyncLockMap[K, V comparable] struct {
	ReadOnly atomic.Bool
	mu       sync.RWMutex
	Map      Map[K, V]
}

// Lock the current map to read-only mode
func (s *SyncLockMap[K, V]) Lock() {
	s.ReadOnly.Store(true)
}

// Unlock the current map
func (s *SyncLockMap[K, V]) Unlock() {
	s.ReadOnly.Store(false)
}

// Set an item with syncronous access
func (s *SyncLockMap[K, V]) Set(k K, v V) error {
	if s.ReadOnly.Load() {
		return ErrReadOnly
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Map[k] = v

	return nil
}

// Get an item with syncronous access
func (s *SyncLockMap[K, V]) Get(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.Map[k]

	return v, ok
}

// Iterate with a callback function synchronously
func (s *SyncLockMap[K, V]) Iterate(f func(k K, v V) error) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k, v := range s.Map {
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

// Clone creates a new SyncLockMap with the same values
func (s *SyncLockMap[K, V]) Clone() *SyncLockMap[K, V] {
	s.mu.Lock()
	defer s.mu.Unlock()

	smap := &SyncLockMap[K, V]{
		ReadOnly: atomic.Bool{},
		mu:       sync.RWMutex{},
		Map:      s.Map.Clone(),
	}
	smap.ReadOnly.Store(s.ReadOnly.Load())
	return smap
}

// Has checks if the current map has the provided key
func (s *SyncLockMap[K, V]) Has(key K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Map.Has(key)
}

// IsEmpty checks if the current map is empty
func (s *SyncLockMap[K, V]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Map.IsEmpty()
}

// GetKeywithValue returns the first key having value
func (s *SyncLockMap[K, V]) GetKeyWithValue(value V) (K, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Map.GetKeyWithValue(value)
}

// Merge the current map with the provided one
func (s *SyncLockMap[K, V]) Merge(n map[K]V) error {
	if s.ReadOnly.Load() {
		return ErrReadOnly
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Map.Merge(n)
	return nil
}

// GetAll returns Copy of the current map
func (s *SyncLockMap[K, V]) GetAll() Map[K, V] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Map.Clone()
}
