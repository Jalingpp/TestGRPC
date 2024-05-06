# TestGRPC
本项目为一个简单的gRPC的使用案例。
## 1 gRPC install
golang-grpc包提供了gRPC相关的代码库，安装命令为：`go get -u google.golang.org/grpc`

安装两个包，用于支持 protobuf 文件的处理：

```
go get -u github.com/golang/protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
```