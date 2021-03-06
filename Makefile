build:
	go build \
		-ldflags "-s -w" \
		-o out/helloWorld \
		./hello_world.go

docker-build:
	docker run -it --rm -v $(PWD):/app -w /app golang:1.13-buster make build
