package main

import (
	"github.com/go-redis/redis"
	"strconv"
	"os"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/transport"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/service"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/pb"
	"net"
	"google.golang.org/grpc"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		//Addr:     getStrParam("REDIS_ADDR"),
		//Password: getStrParam("REDIS_PASSWORD"),
		//DB:       getIntParam("REDIS_DB"),
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	srv, err := service.New(redisClient)
	if err != nil {
		panic(err)
	}

	grpcServer := transport.NewGRPCServer(srv)
	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterShortLinkServiceServer(server, grpcServer)

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
