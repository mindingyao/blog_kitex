package main

import (
	api "blog_kitex/kitex_gen/api/tagservice"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/server"
)

func PrintRequestResponseMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		fmt.Printf("request: %v\n", request)
		err := next(ctx, request, response)
		fmt.Printf("response: %v", response)
		return err
	}
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	svr := api.NewServer(new(TagServiceImpl), server.WithServiceAddr(addr), server.WithMiddleware(PrintRequestResponseMW))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
