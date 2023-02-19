cd bin
GOARCH=386 go build -o calcobot32 ../cmd/main.go
go build -o calcobot ../cmd/main.go
GOARCH=amd64 GOOS=windows go build -o calcobot.exe ../cmd/main.go