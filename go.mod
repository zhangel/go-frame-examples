module github.com/zhangel/go-frame-examples

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/zhangel/go-frame v1.0.1
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506
	google.golang.org/grpc v1.33.1
)

//github.com/golang/protobuf => github.com/golang/protobuf v1.3.3
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
