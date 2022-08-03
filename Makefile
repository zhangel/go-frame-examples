BIN=test-service
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./server/$(BIN) -ldflags="-r ." ./server/main.go
run:build
	./server/$(BIN)  -config.file.type=ini -config.type=file -config.file.path=./config/config_for_dev.ini