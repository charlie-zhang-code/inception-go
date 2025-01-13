# Member

用户基础服务

```bash
gentool -dsn "root:123456@tcp(127.0.0.1:3306)/inception?charset=utf8mb4&parseTime=True&loc=Local" -tables "sys_member" -onlyModel -outPath ./proto/model
```

```bash
goctl rpc protoc ./proto/rpc/member.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style go_zero
```

```bash
go run ./rpc/member.go -f ./rpc/etc/member.yaml
```