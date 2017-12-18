#!/usr/bin/env bash
# 引入 websocket 模块
go get github.com/gorilla/websocket
# 编译订阅器
go build -o subscribe subscribe.go