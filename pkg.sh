cd bin
go build -o calcobot ../cmd/main.go
GOARCH=amd64 GOOS=windows go build -o calcobot.exe ../cmd/main.go