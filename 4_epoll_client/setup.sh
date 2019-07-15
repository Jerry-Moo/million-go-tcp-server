#!/bin/bash
## 这个脚本使用 Docker 在不同的网络命名空间产生多个 client 实例
## 这样才能避免source port 的限制, 在一台机器上才能创建百万的连接
## 
## 用法: ./connect <connections> <number of clients> <server ip>
## Server IP 通常是 Docker gateway IP address, 缺省是 172.17.0.1
## server 端 部署在docker IP address docker 容器名
### ./setup 20000 50 docker.Name

CONNECTIONS=$1
REPLICAS=$2
IP=$3

for ((c=0; c<${REPLICAS}; c++))
do
    docker run -v $(pwd)/client:/client --link million_simple_tcp_server:million_simple_tcp_server  --name million_client_$c -d alpine /client \
    -conn=${CONNECTIONS} -ip=${IP} 
done
