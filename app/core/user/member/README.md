切换目录

```bash
cd E:\MyProjects\crucible_go\app\core\user\member
```

生成Model

```bash
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/crucible_go" -table="sys_user" -dir=./app/core/user/member/proto/model
```

生成API

```bash
goctl api go -api ./proto/api/main.api -dir ./api
```

生成RPC

```bash
goctl rpc protoc ./proto/rpc/main.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc -m
```

运行api

```bash
go run ./api/main.go -f ./api/etc/main.yaml
```

运行rpc

```bash
go run ./rpc/main.go -f ./rpc/etc/main.yaml
```

