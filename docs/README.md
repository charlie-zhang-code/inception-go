# 用户信息基础服务





```bash
docker run -d --name redis-server --network dev-network --restart always -p 6379:6379 redis:latest redis-server --requirepass 123456
```

```bash
docker run -d --name redis-server --network dev-network --restart always -p 6379:6379 redis:latest
```
```bash
docker run -d --name=consul-server --network dev-network --restart always -p 8500:8500 consul:1.15.4 agent -server -bootstrap -ui -node=1 -client='0.0.0.0'
```

```bash
docker run -d --name=consul-server --network dev-network --restart always -p 8500:8500 consul:1.15.4 agent -server -bootstrap -ui -node=default -client='0.0.0.0'
```

```bash
go mod tidy
```

切换目录

```bash
cd E:\MyProjects\crucible_go\app\core\user\member
```

生成Model

```bash
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/crucible_go" -table="sys_user" -dir=./proto/model -c
```

```bash
goctl model mysql ddl --src=proto/sql/schema.sql --dir=proto/model
```

```bash
goctl model mysql ddl --src=proto/sql/schema.sql --dir=proto/model -c
```

生成API

```bash
goctl api go -api ./proto/api/main.api -dir ./api
```

生成RPC

```bash
goctl rpc protoc ./proto/rpc/main.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc -m
```

```bash
goctl rpc protoc ./proto/rpc/main.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc
```

运行api

```bash
go run ./api/main.go -f ./api/etc/main.yaml
```

```yaml
Name: main
Host: 127.0.0.1
Port: 8888

MemberRpc:
  Target: consul://127.0.0.1:8500/user.member.rpc?wait=14s

#Consul:
#  Host: 127.0.0.1:8500 # consul endpoint
#  #  Token: 'f0512db6-76d6-f25e-f344-a98cc3484d42' # consul ACL token (optional)
#  Key: user.member.api
#  Meta:
#    Protocol: grpc
#  Tag:
#    - user.member
#    - api
```

```go
type Config struct {
	rest.RestConf
	MemberRpc zrpc.RpcClientConf
}






type ServiceContext struct {
Config       config.Config
MemberClient service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
return &ServiceContext{
Config:       c,
MemberClient: service.NewService(zrpc.MustNewClient(c.MemberRpc)),
}
}
```

```go
_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"




// register service to consul
_ = consul.RegisterService(c.ListenOn, c.Consul)
```


运行rpc

```bash
go run ./rpc/main.go -f ./rpc/etc/main.yaml
```

```yaml
Name: user.member.rpc
ListenOn: 127.0.0.1:8080
#Etcd:
#  Hosts:
#    - 127.0.0.1:2379
#  Key: user.member.rpc

# mysql: 参数必带
DataSource: root:123456@tcp(127.0.0.1:3306)/crucible_go?parseTime=true

# cache:
Cache:
  - Host: 127.0.0.1:6379

# 替代 etcd:
Consul:
  Host: 127.0.0.1:8500 # consul endpoint
  #  Token: 'f0512db6-76d6-f25e-f344-a98cc3484d42' # consul ACL token (optional)
  Key: user.member.rpc
  Meta:
    Protocol: grpc
  Tag:
    - user.member
    - rpc
```

```go
type Config struct {
	zrpc.RpcServerConf

	// db url:
	DataSource string

	// cache:
	Cache cache.CacheConf

	// 服务注册/发现:
	Consul consul.Conf
}
```

```go
type ServiceContext struct {
	Config config.Config

	Member model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Member:  model.NewSysUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
```



```bash
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"member/proto/model"
	"member/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Member model.SysMemberModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Member: model.NewSysMemberModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}










package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	// db url:
	DataSource string

	// cache:
	Cache cache.CacheConf

	// 服务注册/发现:
	Consul consul.Conf
}






Name: user.member.rpc
ListenOn: 127.0.0.1:8098
#Etcd:
#  Hosts:
#    - 127.0.0.1:2379
#  Key: user.member.rpc

# mysql: 参数必带
DataSource: root:123456@tcp(127.0.0.1:3306)/crucible_go?parseTime=true

# cache:
Cache:
  - Host: 127.0.0.1:6379

# 替代 etcd:
Consul:
  Host: 127.0.0.1:8500 # consul endpoint
  Token: 'f0512db6-76d6-f25e-f344-a98cc3484d42' # consul ACL token (optional)
  Key: user.member.rpc
  Meta:
    Protocol: grpc
  Tag:
    - user.member
    - rpc
```