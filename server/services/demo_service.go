package services

import (
	"github.com/zhangel/go-frame-examples/api"
	"github.com/zhangel/go-frame/server"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DemoService struct {
	server.DefaultProviderImpl
}

func (d *DemoService) GetInfo(ctx context.Context,req *api.GetInfo_Request) (*api.GetInfo_Response, error) {
	dataList := []*api.UserInfo{}
	dataList=append(dataList,&api.UserInfo{Id: int64(1),Name: "tomzhang",Age:10,Address:"beijing"})
	resp:=&api.GetInfo_Response{List: dataList}
	return resp,status.Error(codes.OK,"ok")
}

func (d *DemoService) CreateInfo(ctx context.Context, req *api.CreateInfo_Request)  (*api.Response, error) {
	resp:=&api.Response{Msg:"OK",Code: 0}
	return resp,status.Error(codes.OK,"ok")
}

func (d *DemoService) UpdateInfo(ctx context.Context, req *api.UpdateInfo_Request)  (*api.Response, error) {
	resp:=&api.Response{Msg:"OK",Code: 0}
	return resp,status.Error(codes.OK,"ok")
}

func (d *DemoService) DeleteInfo(ctx context.Context, req *api.DeleteInfo_Request) (*api.Response, error) {
	resp:=&api.Response{Msg:"OK",Code: 0}
	return resp,status.Error(codes.OK,"ok")
}
