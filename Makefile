build:
	go build -o bin/party-calc ./cmd/main.go
run:
	go run ./cmd/main.go -confpath=./cmd
exe:
	./bin/party-calc
	
docker-build:
	docker build --tag party-calc .
docker-run:
	docker run -p 8080:8080 -e DB_HOST=127.0.0.1 -e DB_PORT=5432 -e DB_USER="postgres" -e DB_PASSWORD="password" -e DB_NAME="partycalc"  party-calc

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out

swag:
	swag fmt
	swag init -g ./internal/server/http/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html
