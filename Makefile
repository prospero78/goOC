run.app:
	clear
	go fmt ./...
	go build -o ./bin/oc ./cmd/gooc/main.go
	./bin/oc ./bin/Main.o7
mod:
	clear
	go fmt ./...
	go mod tidy
	go mod vendor
lint:
	clear
	go fmt ./...
	golangci-lint run
	gocritic check
	gocyclo -top 5 ./internal
run.doc:
	godoc ./cmd/gooc
fmt:
	go fmt ./...
