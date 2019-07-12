#!/bin/bash
## 移除所有测试客户端, 名称以 million_ 开头的容器都会被删除
## be carefully

docker rm -vf $(docker ps -a --format '{{.ID}} {{.Names}}'|grep 'million_'|awk '{print $1}' )