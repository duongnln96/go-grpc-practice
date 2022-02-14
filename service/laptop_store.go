package service

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	// Search searching laptop
	Search(ctx context.Context, filter *pcbook.Filter, found func(laptop *pcbook.Laptop) error) error
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

func (store *inMemoryLaptopStore) Search(ctx context.Context, filter *pcbook.Filter, found func(laptop *pcbook.Laptop) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Printf("context is cancelled \n")
			return nil
		}

		if isQualified(filter, laptop) {
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func isQualified(filter *pcbook.Filter, laptop *pcbook.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(memory *pcbook.Memory) uint64 {
	value := memory.GetValue()

	switch memory.GetUnit() {
	case pcbook.Memory_BIT:
		return value
	case pcbook.Memory_BYTE:
		return value << 3 // 8 = 2^3
	case pcbook.Memory_KILOBYTE:
		return value << 13 // 1024 * 8 = 2^10 * 2^3 = 2^13
	case pcbook.Memory_MEGABYTE:
		return value << 23
	case pcbook.Memory_GIGABYTE:
		return value << 33
	case pcbook.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}
