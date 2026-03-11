APP := gin-websocket-boilerplate

.PHONY: build run test lint docker-build docker-run

build:
	go build -o $(APP) .

run: build
	./$(APP)

test:
	go test ./...

lint:
	go vet ./...

docker-build:
	docker build -t $(APP) .

docker-run:
	docker run --rm -p 8000:8000 $(APP)
