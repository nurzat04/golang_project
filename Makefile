.PHONY: all build run-short run-redirect run-stats clean docker

all: build

build:
	go build -o bin/shorten ./cmd/shorten
	go build -o bin/redirect ./cmd/redirect
	go build -o bin/stats ./cmd/stats

run-short:
	PORT=8080 go run ./cmd/shorten

run-redirect:
	PORT=8081 go run ./cmd/redirect

run-stats:
	PORT=8082 go run ./cmd/stats

clean:
	rm -rf bin

docker:
	docker build -t shorty .

swagger-ui:
	docker run --rm -p 8888:8080 -e SWAGGER_JSON=/spec/swagger.yaml -v $(PWD)/docs/swagger.yaml:/spec/swagger.yaml swaggerapi/swagger-ui