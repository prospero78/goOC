run.app:
	clear
	go fmt ./...
	go build -o ./bin/oc ./cmd/gooc/main.go
	cd ./bin && \
	export DEBUG=true && \
	./oc ./src/Main.o7
run.gui:
	clear
	go fmt ./...
	go build -o ./bin/guioc ./cmd/guioc/main.go
	cd ./bin && \
	export DEBUG=true && \
	./guioc ./src/Main.o7
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
