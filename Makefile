build:
	sh docker/build.sh

local-run:
	docker run -it -p 8080:8080 golang-echo-sample:latest