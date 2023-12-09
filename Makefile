.DEFAULT_GOAL := docker

start:
	go run main.go
	
cli:
	go build ./packages/cli/main.go

docker:
	docker build . -t kletskovg/accounting:server --file=Dockerfile.server
docker_run:
	docker run kletskovg/accounting:server -p "8080:8080"