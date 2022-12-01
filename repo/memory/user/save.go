package user

import (
	errCode "github.com/duongnln96/go-grpc-practice/constants/error_code"
	"github.com/duongnln96/go-grpc-practice/model"
)

// Save saves a user to the store
func (store *inMemoryUserStore) Save(user *model.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errCode.ErrAlreadyExists
	}

	store.users[user.Username] = user.Clone()
	return nil
}
