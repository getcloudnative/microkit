// Package memory provides a service that implements a memory storage.
package memory

import (
	"sync"

	microerror "github.com/giantswarm/microkit/error"
)

// Config represents the configuration used to create a memory service.
type Config struct {
}

// DefaultConfig provides a default configuration to create a new memory
// service by best effort.
func DefaultConfig() Config {
	return Config{}
}

// New creates a new configured memory service.
func New(config Config) (*Service, error) {
	newService := &Service{
		storage: map[string]string{},
		mutex:   sync.Mutex{},
	}

	return newService, nil
}

// Service is the memory service.
type Service struct {
	// Internals.
	storage map[string]string
	mutex   sync.Mutex
}

func (s *Service) Create(key, value string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.storage[key] = value

	return nil
}

func (s *Service) Delete(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.storage, key)

	return nil
}

func (s *Service) Exists(key string) (bool, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.storage[key]

	return ok, nil
}

func (s *Service) Search(key string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	value, ok := s.storage[key]
	if ok {
		return value, nil
	}

	return "", microerror.MaskAnyf(keyNotFoundError, key)
}