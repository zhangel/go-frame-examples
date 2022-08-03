BIN=test-client
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(BIN) -ldflags="-r ." main.go
run:build
	./$(BIN)