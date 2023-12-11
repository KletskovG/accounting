.DEFAULT_GOAL := build

start:
	go run main.go

build:
	alias cli="/home/kletskov02/.accounting/cli"
	go build
	cp ./cli ~/.accounting/cli
	cp ./config.env ~/.accounting/config.env
	rm ./cli
