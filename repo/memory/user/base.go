package user

import (
	"sync"

	"github.com/duongnln96/go-grpc-practice/model"
)

// UserStore is an interface to store users
type UserStore interface {
	// Save saves a user to the store
	Save(user *model.User) error
	// Find finds a user by username
	Find(username string) (*model.User, error)
	// SeedUser creates some users
	SeedUser() error
}

// inMemoryUserStore stores users in memory
type inMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*model.User
}

// NewInMemoryUserStore returns a new in-memory user store
func NewInMemoryUserStore() *inMemoryUserStore {
	return &inMemoryUserStore{
		users: make(map[string]*model.User),
	}
}

func (s *inMemoryUserStore) SeedUser() error {
	err := s.createUser("admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return s.createUser("user1", "secret", "user")
}

func (s *inMemoryUserStore) createUser(username, password, role string) error {
	user, err := model.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return s.Save(user)
}
