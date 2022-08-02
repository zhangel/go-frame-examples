package main

import (
	"context"
	"github.com/zhangel/go-frame"
	"github.com/zhangel/go-frame/log"
)

func main() {
	defer framework.Init()()
	logger:=log.WithContext(context.Background())
	logger.Info("test info")
}


