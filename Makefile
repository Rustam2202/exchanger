build:
	go build -o bin/exchanger ./cmd/exchanges/main.go
run:
	go run ./cmd/exchanges/main.go -confpath=./
exe:
	./bin/exchanger
	
test:
	go test ./...

swag:
	swag fmt
	swag init -g ./internal/server/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html
