build:
	go build -mod=vendor -o ./build/poll-api ./cmd/poll-api/main.go

build:docker:
	docker build --file=./build/docker/Dockerfile .