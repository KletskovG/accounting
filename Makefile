lint:
	bash ./tools/lint.sh

test:
	bash ./tools/test.sh || exit 1

db_testing:
	sudo docker run -p 27017:27017 mongo:latest