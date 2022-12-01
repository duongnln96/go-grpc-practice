package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Role           string `json:"role"`
}

func (m *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(password))
	return err == nil
}

func (m *User) Clone() *User {
	return &User{
		Username:       m.Username,
		HashedPassword: m.HashedPassword,
		Role:           m.Role,
	}
}

func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}

	return user, nil
}
