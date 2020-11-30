run.app:
	clear
	export GOMODULE=on && \
	go build -o ./bin/oc ./cmd/gooc/main.go
	./bin/oc ./bin/Hello.o7
fmt:
	go fmt ./...
