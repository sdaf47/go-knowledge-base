package main

import (
	"github.com/go-redis/redis"
	"strconv"
	"os"
	"github.com/sdaf47/go-knowledge-base/shortlink/linkshort/transport"
	"github.com/sdaf47/go-knowledge-base/shortlink/linkshort/service"
	"github.com/sdaf47/go-knowledge-base/shortlink/linkshort/pb"
	"net"
	"google.golang.org/grpc"
	"fmt"
)

func main() {
	fmt.Println("Short link server starting . . . ")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     getStrParam("REDIS_ADDR"),
		Password: getStrParam("REDIS_PASSWORD"),
		DB:       getIntParam("REDIS_DB"),
	})

	err := redisClient.Keys("**").Err()
	if err != nil {
		panic(err)
	}

	srv, err := service.New(redisClient)
	if err != nil {
		panic(err)
	}

	grpcServer := transport.NewGRPCServer(srv)
	grpcListener, err := net.Listen("tcp", getStrParam("SHORT_LINK_ADDR"))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterShortLinkServiceServer(server, grpcServer)

	fmt.Println("Short link server started.")
	server.Serve(grpcListener)
}

func getIntParam(key string) (val int) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}

	return
}

func getStrParam(key string) (val string) {
	val = os.Getenv(key)
	return
}
