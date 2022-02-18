build:
	sh docker/build.sh

test:
	docker run -it -p 8080:8080 golang-echo-sample:latest