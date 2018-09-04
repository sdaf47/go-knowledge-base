package service

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/database"
	"github.com/minio/highwayhash"
	"github.com/go-redis/redis"
	"encoding/hex"
	"fmt"
)

const ShortLinkSize = 10

var highWayKey []byte

type ShortLinkService interface {
	Decode(code string) (address string, err error)
	Encode(link string) (code string, err error)
}

type service struct {
	linkRepo database.LinkRepo
}

func New(db *redis.Client) (s *service, err error) {
	highWayKey, err = hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	if err != nil {
		return
	}

	s = &service{
		linkRepo: database.New(db),
	}

	return
}

func (s *service) Decode(code string) (link string, err error) {
	return s.linkRepo.Link(code)
}

func (s *service) Encode(link string) (code string, err error) {
	var res string
	var i = 0
	for {
		code, err = hash(link + string(i))
		if err != nil {
			return
		}

		res, err = s.linkRepo.Link(code)
		if res == link {
			return
		}
		if res == "" {
			break
		}
		i++
	}

	err = s.linkRepo.Create(code, link)

	return
}

func hash(address string) (code string, err error) {
	h := highwayhash.Sum128([]byte(address), highWayKey)
	code = fmt.Sprintf("%x", h[:ShortLinkSize])

	return
}
