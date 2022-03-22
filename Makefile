build:
	go build -o main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 go build -o main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o main-windows-386 main.go

tidy:
	go mod tidy

init:
	go mod init $(module-name)

