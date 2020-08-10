package database

import (
	"context"
	"github.com/go-redis/redis"
)

const RedisExpire = 0

type LinkRepo interface {
	Create(code, link string) (err error)
	Link(code string) (link string, err error)
}

type linkRepo struct {
	db *redis.Client
}

func New(db *redis.Client) LinkRepo {
	return &linkRepo{
		db: db,
	}
}

func (repo *linkRepo) Create(code, link string) (err error) {
	_, err = repo.db.Set(context.Background(), code, link, RedisExpire).Result()
	if err != nil {
		return
	}

	return
}

func (repo *linkRepo) Link(code string) (link string, err error) {
	return repo.db.Get(context.Background(), code).Result()
}
