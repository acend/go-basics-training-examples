package api

import (
	"fmt"
	"sync"
)

type UserStore interface {
	List() ([]User, error)
	Create(user *User) error
	Get(name string) (*User, error)
	Delete(name string) error
}

var _ UserStore = &MemoryUserStore{}

type MemoryUserStore struct {
	users map[string]User
	mu    *sync.Mutex
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: map[string]User{},
		mu:    &sync.Mutex{},
	}
}

// Create implements UserStore
func (us *MemoryUserStore) Create(user *User) error {
	us.mu.Lock()
	us.users[user.Name] = *user
	us.mu.Unlock()
	return nil
}

// Delete implements UserStore
func (us *MemoryUserStore) Delete(name string) error {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, ok := us.users[name]
	if !ok {
		return fmt.Errorf("not found")
	}
	delete(us.users, name)
	return nil
}

// Get implements UserStore
func (us *MemoryUserStore) Get(name string) (*User, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[name]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &user, nil
}

// List implements UserStore
func (us *MemoryUserStore) List() ([]User, error) {
	us.mu.Lock()

	users := []User{}
	for _, user := range us.users {
		users = append(users, user)
	}

	us.mu.Unlock()
	return users, nil
}
