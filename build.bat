echo off
cls
go fmt ./...
go test ./...
go build main.go
move main.exe .\bin\main.exe
.\bin\strip -s .\bin\main.exe
.\bin\upx .\bin\main.exe
