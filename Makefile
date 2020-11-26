run.app:
	go build -mod=vendor -o ./bin/oc ./cmd/gooc/main.go
	./bin/oc
