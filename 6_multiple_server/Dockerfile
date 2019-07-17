FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

# copy ./server to /app
ADD server /app/server

ENTRYPOINT ["./server"]
CMD ["./server"]