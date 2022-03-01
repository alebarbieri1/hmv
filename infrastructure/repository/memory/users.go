package memory

import (
	"flavioltonon/hmv/domain/entity"
	"sync"
	"time"
)

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(e *entity.User) *User {
	return &User{
		ID:        e.ID,
		Username:  e.Username,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (u *User) toEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UsersRepository struct {
	users map[string]*User
	mu    sync.RWMutex
}

func NewUsersRepository() (*UsersRepository, error) {
	return &UsersRepository{users: make(map[string]*User)}, nil
}

func (r *UsersRepository) CreateUser(user *entity.User) error {
	r.mu.Lock()
	r.users[user.ID] = NewUser(user)
	r.mu.Unlock()
	return nil
}

func (r *UsersRepository) FindUserByID(userID string) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return user.toEntity(), nil
}

func (r *UsersRepository) FindUserByUsername(username string) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Username == username {
			return user.toEntity(), nil
		}
	}

	return nil, entity.ErrNotFound
}

func (r *UsersRepository) UpdateUser(user *entity.User) error {
	r.mu.Lock()
	r.users[user.ID] = NewUser(user)
	r.mu.Unlock()
	return nil
}
