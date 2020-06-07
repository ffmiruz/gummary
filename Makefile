build:
	mkdir -p functions
	go get ./...
	go build -o functions/api nfcmd/main.go
