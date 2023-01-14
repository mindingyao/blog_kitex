package main

import (
	"blog_kitex/kitex_gen/api"
	"blog_kitex/kitex_gen/api/tagservice"
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := tagservice.NewClient("example", client.WithHostPorts("127.0.0.1:8000"))
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GetTagListRequest{Name: "golang", Page: 1, PageSize: 2}
	resp, err := c.GetTagList(context.Background(), req, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
