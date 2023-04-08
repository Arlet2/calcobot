GOARCH=386      GOOS=linux      go build -o bin/calcobot32      cmd/main.go
GOARCH=amd64    GOOS=linux      go build -o bin/calcobot        cmd/main.go
GOARCH=amd64    GOOS=windows    go build -o bin/calcobot.exe    cmd/main.go