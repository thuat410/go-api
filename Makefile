MAIN_PATH=./cmd/server
APP_NAME=go-api
run:
	go run $(MAIN_PATH)/main.go

build:
	go build -o $(APP_NAME) $(MAIN_PATH)/main.go

install:
	go install $(MAIN_PATH)
