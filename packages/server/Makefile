.DEFAULT_GOAL := start

start:
	go run main.go

build:
	go build

docker:
	docker build . -t kletskovg/test
docker_run:
	docker run kletskovg/test