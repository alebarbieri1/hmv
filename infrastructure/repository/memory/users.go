package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

// User is a representation of entity.User in the repository
type User struct {
	ID          string
	Username    string
	Password    string
	ProfileKind valueobject.ProfileKind
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewUser creates a new User
func NewUser(e *entity.User) *User {
	return &User{
		ID:          e.ID,
		Username:    e.Username,
		Password:    e.Password,
		ProfileKind: e.ProfileKind,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

// toEntity transforms an User into a entity.User
func (u *User) toEntity() *entity.User {
	return &entity.User{
		ID:          u.ID,
		Username:    u.Username,
		Password:    u.Password,
		ProfileKind: u.ProfileKind,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// UsersRepository is a repository for entity.User entities
type UsersRepository struct {
	users map[string]*User
	mu    sync.RWMutex
}

// NewUsersRepository creates a new UsersRepository
func NewUsersRepository() (*UsersRepository, error) {
	return &UsersRepository{users: make(map[string]*User)}, nil
}

// CreateUser stores an entity.User into the repository. If an User with the same ID already
// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
func (r *UsersRepository) CreateUser(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return entity.ErrDuplicatedEntry
	}

	r.users[user.ID] = NewUser(user)
	return nil
}

// FindUserByID returns an entity.User identified by a given userID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *UsersRepository) FindUserByID(userID string) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return user.toEntity(), nil
}

// FindUserByUsername returns an entity.User identified by a given username. If no entities are found,
// entity.ErrNotFound should be returned instead.
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

// ListUsers returns all the entity.User of the repository
func (r *UsersRepository) ListUsers() ([]*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*entity.User, 0, len(r.users))

	for _, user := range r.users {
		users = append(users, user.toEntity())
	}

	return users, nil
}

// UpdateUser updates an User in the repository. If no entities with the same ID as the input entity.User are found,
// entity.ErrNotFound should be returned instead.
func (r *UsersRepository) UpdateUser(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return entity.ErrNotFound
	}

	r.users[user.ID] = NewUser(user)
	return nil
}
