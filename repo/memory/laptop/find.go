package laptop

import (
	"github.com/duongnln96/go-grpc-practice/pb"
)

func (store *inMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deepCopy(laptop)
}
