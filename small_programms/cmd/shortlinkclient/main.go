package main

import (
	"strconv"
	"os"
	"google.golang.org/grpc"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/transport"
	"fmt"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	svc := transport.NewGRPCClient(conn)

	code, err := svc.Encode("https://github.com/go-kit/kit/blob/master/examples/addsvc/cmd/addcli/addcli.go")
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	fmt.Println("code", code)

	link, err := svc.Decode(code)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	fmt.Println("link", link)
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
