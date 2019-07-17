build:
	GOOS=linux GOARCH=amd64 go build --tags "linux" -o server
	docker build -t million-server_workpool .

run:
	docker run --name million_simple_tcp_server  million-server_workpool -c 10 