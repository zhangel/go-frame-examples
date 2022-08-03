package services

import (
	"fmt"
	"github.com/zhangel/go-frame-examples/api"
	"github.com/zhangel/go-frame/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
	"context"
)

type TestApiServer struct {
	server.DefaultProviderImpl
}

func (s *TestApiServer) Unary(ctx context.Context, req *api.Api_Request) (*api.Api_Response, error) {
	logger := s.Logger().WithContext(ctx)
	logger.Tracef("Unary, request = %+v", req)

	grpc.SetHeader(ctx, metadata.New(map[string]string{
		"Header1": "Value1",
	}))
	//time.Sleep(10* time.Second)

	return &api.Api_Response{Data: "Unary response", Code: 0}, nil
}

func (s *TestApiServer) ServerSideStreaming(req *api.Api_Request, streamer api.TestApi_ServerSideStreamingServer) error {
	logger := s.Logger().WithContext(streamer.Context())
	logger.Tracef("ServerSideStreaming, request = %+v", req)

	metadata.NewOutgoingContext(streamer.Context(), metadata.New(map[string]string{"TEST": "TEST"}))

	md := metadata.New(map[string]string{
		"Header1": "Value1",
		"Header2": "Value1",
		"Header3": "Value1",
		"Header4": "Value1",
		"Header5": "Value1",
	})
	err := streamer.SetHeader(md)
	md = metadata.New(map[string]string{
		"Header6": "Value1",
		"Header7": "Value1",
		"Header8": "Value1",
		"Header9": "Value1",
	})
	streamer.SendHeader(md)
	fmt.Println("SetHeader got err:", err)

	md = metadata.New(map[string]string{
		"Trailer":  "Value1",
		"Trailer1": "Value1",
		"Trailer2": "Value1",
		"Trailer3": "Value1",
	})
	grpc.SetTrailer(streamer.Context(), md)
	md = metadata.New(map[string]string{
		"Trailer4": "Value1",
		"Trailer5": "Value1",
	})
	streamer.SetTrailer(md)
	for i := 0; i < 10; i++ {
		logger.Tracef("ServerSideStreaming, Send response for %d", i)
		err := streamer.Send(&api.Api_Response{Data: fmt.Sprintf("ServerSideStreaming response for %d", i), Code: int32(i)})
		if err != nil {
			logger.Errorf("ServerSideStreaming server::Send failed, err = %v", err)
			break
		}
		time.Sleep(time.Second)
	}

	return fmt.Errorf("Hello error !")
	//return nil
}

func (s *TestApiServer) ClientSideStreaming(streamer api.TestApi_ClientSideStreamingServer) error {
	logger := s.Logger().WithContext(streamer.Context())

	for {
		req, err := streamer.Recv()
		if err != nil {
			if err != io.EOF {
				logger.Error(err)
			}
			break
		}
		logger.Trace(req, err)
	}

	err := streamer.SendAndClose(&api.Api_Response{Data: fmt.Sprintf("ClientSideStreaming response"), Code: 0})
	if err != nil {
		logger.Errorf("ClientSideStream server::SendAndClose failed, err = %v", err)
	}

	return nil
}

func (s *TestApiServer) BidiSideStreaming(streamer api.TestApi_BidiSideStreamingServer) error {
	logger := s.Logger().WithContext(streamer.Context())

	for {
		req, err := streamer.Recv()
		if err != nil {
			if err != io.EOF {
				logger.Error(err)
			}
			break
		}
		logger.Trace(req, err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		logger.Tracef("BidiStreaming, Send response for %d", i)
		err := streamer.Send(&api.Api_Response{Data: fmt.Sprintf("BidiStreaming response for %d", i), Code: int32(i)})
		if err != nil {
			if err != io.EOF {
				logger.Error(err)
			}
			break
		}
		logger.Tracef("BidiStreaming, Send response for %d successful", i)
	}

	fmt.Println("BidiSideStreaming Exit")
	return nil
}

