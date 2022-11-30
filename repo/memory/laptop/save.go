package laptop

import (
	errCode "github.com/duongnln96/go-grpc-practice/constants/error_code"
	"github.com/duongnln96/go-grpc-practice/pb"
)

// Save saves the laptop to the store
func (store *inMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return errCode.ErrAlreadyExists
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}
