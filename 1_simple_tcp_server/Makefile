build:
	GOOS=linux GOARCH=amd64 go build -o server server.go
	docker build -t million-simple_tcp_server .

run:
	docker run --name million_simple_tcp_server million-simple_tcp_server