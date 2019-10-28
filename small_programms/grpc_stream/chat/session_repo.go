package chat

import (
	"fmt"
	"sync"
	"time"
)

type SessionRepo interface {
	Authorize(login, password string) (token string, err error)
	IsAuthorized(token string) bool
	Unauthorized(token string) error
}

type sessionRepo struct {
	mu   sync.Mutex
	mock map[string]bool
}

func NewSessionRepo() SessionRepo {
	return &sessionRepo{
		mock: map[string]bool{},
	}
}

func (repo *sessionRepo) Authorize(login, password string) (token string, err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	token = fmt.Sprintf("%X", time.Now().UnixNano())
	repo.mock[token] = true

	return
}

func (repo *sessionRepo) IsAuthorized(token string) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.mock[token]
}

func (repo *sessionRepo) Unauthorized(token string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.mock, token)

	return nil
}
