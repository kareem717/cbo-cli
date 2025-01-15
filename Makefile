dev:
	@go build -tags dev -o bin/k7-cbo main.go

build:
	@go build -o bin/k7-cbo main.go

run: build
	@./bin/k7-cbo $(filter-out $@,$(MAKECMDGOALS))

install:
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download

test:
	@go test ./...
