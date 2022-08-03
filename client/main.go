package main

import (
	"context"
	"fmt"
	framework "github.com/zhangel/go-frame"
	"github.com/zhangel/go-frame-examples/api"
	"github.com/zhangel/go-frame/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

var (
	edge_addr = "172.22.118.198:9988"
)

func dial() *grpc.ClientConn {
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	conn,err:=grpc.DialContext(ctx,edge_addr,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial fail,error=%+v",err)
	}
	return conn
}

//test code main
func main() {
	defer framework.Init()()
	conn:=dial()
	client:=api.NewInfoClient(conn)
	TestInfo(client)
	apiClient:=api.NewTestApiClient(conn)
	TestApi(apiClient)
}

func TestApi(client api.TestApiClient) {
	TestUnary(client)
	TestServerSideStreaming(client)
	TestServerSideStreamingWithoutRecv(client)
	TestClientSideStreaming(client)
	TestBidiStreaming(client)
}

func TestInfo(client api.InfoClient) {
	TestGetInfo(client)
	TestCreateInfo(client)
	TestUpdateInfo(client)
	TestDeleteInfo(client)
}

func TestGetInfo(c api.InfoClient) {
	req:=&api.GetInfo_Request{}
	resp,err:=c.GetInfo(context.Background(),req)
	log.Trace(resp,err)
}

func TestCreateInfo(c api.InfoClient) {
	req:=&api.CreateInfo_Request{Name: "cameron zhang",Address: "beijing",Age: 20}
	resp,err:=c.CreateInfo(context.Background(),req)
	log.Trace(resp,err)
}

func TestUpdateInfo(c api.InfoClient) {
	req:=&api.UpdateInfo_Request{Name: "cameron liu",Address: "xian",Age: 22}
	resp,err:=c.UpdateInfo(context.Background(),req)
	log.Trace(resp,err)
}

func TestDeleteInfo(c api.InfoClient) {
	req:=&api.DeleteInfo_Request{Id: 1}
	resp,err:=c.DeleteInfo(context.Background(),req)
	log.Trace(resp,err)
}

func TestUnary(client api.TestApiClient) {
	fmt.Println("UnaryTest")
	defer fmt.Println("UnaryTest DONE")

	header := metadata.New(nil)
	trailer := metadata.New(nil)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := client.Unary(
		ctx,
		&api.Api_Request{Data: "Unary REQ"},
	)

	fmt.Println("Header:", header)
	fmt.Println("Trailer:", trailer)

	log.Trace(resp, err)
}

func TestServerSideStreaming(client api.TestApiClient) {
	fmt.Println("ServerSideStreamingTest")
	defer fmt.Println("ServerSideStreamingTest DONE")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	header := metadata.New(nil)
	trailer := metadata.New(nil)

	streamer, err := client.ServerSideStreaming(ctx, &api.Api_Request{Data: "ServerSideStreaming REQ"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}

	for {
		if req, err := streamer.Recv(); err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			break
		} else {
			log.Trace(req)
			fmt.Println("Trailer:", streamer.Trailer())
		}
	}

	fmt.Println("Header:", header)
	fmt.Println("Trailer:", trailer)
}

func TestServerSideStreamingWithoutRecv(client api.TestApiClient) {
	fmt.Println("ServerSideStreamingWithoutRecvTest")
	defer fmt.Println("ServerSideStreamingWithoutRecvTest DONE")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.ServerSideStreaming(ctx, &api.Api_Request{Data: "ServerSideStreaming REQ"})
	if err != nil {
		log.Fatal(err)
	}

	return
}

func TestClientSideStreaming(client api.TestApiClient) {
	fmt.Println("ClientSideStreamingTest")
	defer fmt.Println("ClientSideStreamingTest DONE")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	streamer, err := client.ClientSideStreaming(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		log.Tracef("ClientSideStreaming client::Send, client::Send REQ for %d", i)
		err := streamer.Send(&api.Api_Request{Data: fmt.Sprintf("ClientSideStreaming REQ for %d", i)})
		if err != nil {
			log.Errorf("ClientSideStreaming client::Send failed, err = %v", err)
			break
		}
	}

	resp, err := streamer.CloseAndRecv()
	log.Trace(resp, err)
}

func TestBidiStreaming(client api.TestApiClient) {
	fmt.Println("BidiStreamingTest")
	defer fmt.Println("BidiStreamingTest DONE")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamer, err := client.BidiSideStreaming(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		log.Tracef("BidiStreaming client::Send, client::Send REQ for %d", i)
		err := streamer.Send(&api.Api_Request{Data: fmt.Sprintf("BidiStreaming REQ for %d", i)})
		if err != nil {
			log.Errorf("BidiStreaming client::Send failed, err = %v", err)
			break
		}
	}

	log.Tracef("BidiStreaming client::CloseSend")
	err = streamer.CloseSend()
	if err != nil {
		log.Errorf("BidiStreaming client::CloseSend failed, err = %v", err)
	}

	for {
		if resp, err := streamer.Recv(); err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			break
		} else {
			log.Trace(resp)
		}
	}
}


