package main

import (
	"context"
	framework "github.com/zhangel/go-frame"
	"github.com/zhangel/go-frame-examples/server/services"
	"github.com/zhangel/go-frame/log"
	"github.com/zhangel/go-frame/server"
	"github.com/zhangel/go-frame-examples/api"
)

func main() {
	defer framework.Init()()
	logger:=log.WithContext(context.Background())
	var err error
	err=server.RegisterService(api.RegisterTestApiServer,&services.TestApiServer{})
	if err != nil {
		logger.Fatalf("register service fail,error=%+v",err)
	}
	err=server.RegisterService(api.RegisterInfoServer,&services.DemoService{})
	if err != nil {
		logger.Fatalf("register service fail,error=%+v",err)
	}
	err=server.Run()
	if err != nil {
		logger.Fatalf("run service fail,error=%+v",err)
	}
}


