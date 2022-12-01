package user

import "github.com/duongnln96/go-grpc-practice/model"

// Find finds a user by username
func (store *inMemoryUserStore) Find(username string) (*model.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
