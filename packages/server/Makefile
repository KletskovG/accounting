.DEFAULT_GOAL := start

start:
	go run main.go

build:
	go build

test:
	ACC_DB_URL="mongodb://localhost:27017/" ACC_DB_NAME=test ACC_DB_COLLECTION=test  go test  ./...

test_mongo:
	sudo docker run -p 27017:27017 mongo:latest