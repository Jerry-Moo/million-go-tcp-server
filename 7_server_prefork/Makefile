build:
	GOOS=linux GOARCH=amd64 go build --tags "linux" -o server
	docker build -t million-epoll_server_prefork .

run:
	docker run --name million_simple_tcp_server  -d million-epoll_server_prefork -c 10 -prefork