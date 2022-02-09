package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/duongnln96/go-grpc-practice/pb/pcbook"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pcbook.Laptop) error
	// Find finds a laptop by ID
	Find(id string) (*pcbook.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type inMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pcbook.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() LaptopStore {
	return &inMemoryLaptopStore{
		data: make(map[string]*pcbook.Laptop),
	}
}

// Save saves the laptop to the store
func (store *inMemoryLaptopStore) Save(laptop *pcbook.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(laptop *pcbook.Laptop) (*pcbook.Laptop, error) {
	other := &pcbook.Laptop{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}

func (store *inMemoryLaptopStore) Find(id string) (*pcbook.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deepCopy(laptop)
}
