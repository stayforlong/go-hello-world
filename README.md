# go-hello-world
Simple HTTP Hello World in Golang

# Build

### Requirements
- docker installed and running

### How to build
From the root of this repository run `make docker-build` and the binary will be located in `./out/` directory (the binary will be compiled for `linux amd64` platform only).

# Run
Just execute the binary and the webserver will be available at port `80`

### Endpoints

- `/` -> Hello World
- `/health` -> health check
