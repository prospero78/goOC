
echo Сборка oc
cd ./src/
go build -mod=vendor -o ../bin/oc ./cmd/main.go
cd ../
