.DEFAULT_GOAL := docker

start:
	go run main.go
	
cli:
	go build ./packages/cli/main.go

docker:
	docker build . -t kletskovg/accounting:server -f deployment/Dockerfile.server
docker_run:
	docker run  -p "8080:8080" kletskovg/accounting:server