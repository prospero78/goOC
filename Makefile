run.app:
	clear
	export GOMODULE=on && \
	go build -o ./bin/oc ./cmd/gooc/main.go
	./bin/oc ./bin/Main.o7
run.doc:
	godoc ./cmd/gooc
fmt:
	go fmt ./...
