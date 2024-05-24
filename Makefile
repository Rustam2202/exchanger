build:
	go build -o bin/exchanger ./cmd/exchanges/main.go
run:
	go run ./cmd/exchanges/main.go -confpath=./
exe:
	./bin/exchanger
	
docker-build:
	docker build --tag exchanger .
docker-run:
	docker run -p 8080:8080 exchanger

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover:
	make test
	go tool cover -html=coverage.out

swag:
	swag fmt
	swag init -g ./internal/server/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html
