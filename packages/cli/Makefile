.DEFAULT_GOAL := build

start:
	go run main.go

build:
	alias cli="/home/kletskov02/Software/accounting/packages/cli"
	go build

watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' build 


docker:
	docker build . -t kletskovg/test
docker_run:
	docker run kletskovg/test