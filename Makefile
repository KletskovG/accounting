.DEFAULT_GOAL := start

start:
	go run main.go

docker:
	docker build . -t kletskovg/test
docker_run:
	docker run kletskovg/test