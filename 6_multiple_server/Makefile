build:
	GOOS=linux GOARCH=amd64 go build -tags "linux" -o server
	docker build -t million-multiple-server .

run:
	docker run --name million_simple_tcp_server million-multiple-server