all: docker-up

build:
	CGO_ENABLED="0" GOOS=linux GOARCH=amd64 go build -o main

docker: build
	docker build -t "charlesbases/swagger-ui:latest" .
	@rm -rf main

docker-up: docker
	docker-compose up -d
