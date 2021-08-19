
# grpc

## 1. echo 服务

1. 创建 echo.proto 并定义请求响应消息以及服务描述

message 用于序列化消息

service 用于 grpc

2. gen.sh 脚本

```sh
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       protobuf/echo.proto
```

需要安装 protobuf（包含protoc命令）、protoc-gen-go（消息的序列化反序列化）、protoc-gen-go-grpc（grpc协议）

3. 实现 server

4. client 调用

## 2. TODO：一个更实用的例子
