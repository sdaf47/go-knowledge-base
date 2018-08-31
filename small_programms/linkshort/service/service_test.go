package service

import (
	"testing"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var srv ShortLinkService

func TestMain(m *testing.M) {
	client := redis.NewClient(&redis.Options{
		//Addr:     os.Getenv("REDIS_ADDR"),
		//Password: os.Getenv("REDIS_PASSWORD"),
		//DB:       getIntParam("REDIS_DB"),
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var err error
	srv, err = New(client)
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestService_GetLinkByShort(t *testing.T) {
	testLinks := []string{
		"https://domain.ru/aasdaasdasd/asdasdas/dasdasd/?asdasd=asdasdasd",
		"https://domain.ru/aasdaasdasd/asdasdas/dasdasd/?asdasd=13123123123",
		"http://domain.ru/aasdaasdasd/asdasdas/?aaa=qqq1&asdasd=asdasdasd",
		"http://domain.ru/aasdaasdasd/asdasdas/?aaa=qqq1&ttt=12&asdasd=asdasdasd",
		"http://domain.ru/aasdaasdasd/?aaa=qqq1&ttt=12&asdasd=asdasdasd",
		"https://test.domain.ru/aasdaasdasd/asdasdas/?aaa=qqq1&ttt=12&asdasd=asdasdasd",
		"https://site.domain.ru/aasdaasdasd/asdasdas/dasdasd/?asdasd=asdasdasd",
	}

	var code string
	var address string
	var err error
	for _, link := range testLinks {
		code, err = srv.Encode(link)
		if err != nil {
			t.Fatal(err)
		}

		address, err = srv.Decode(code)
		if address != link {
			t.Fatalf("Link does not equil: %s != %s\n", address, link)
		}

	}
}

func getIntParam(key string) (val int) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}

	return
}
