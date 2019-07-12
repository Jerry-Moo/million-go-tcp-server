FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

# 将当前目录下 ./server 拷贝到 /app 目录下
ADD server /app/server

ENTRYPOINT ["./server"]
CMD ["./server"]