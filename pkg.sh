cd bin
go build ../cmd/main.go
GOARCH=amd64 GOOS=windows go build ../cmd/main.go