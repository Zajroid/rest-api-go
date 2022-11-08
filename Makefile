BINARY_NAME=main.out

run:
	go run ./cmd/main/app.go

deps:
	go get github.com/julienschmidt/httprouter