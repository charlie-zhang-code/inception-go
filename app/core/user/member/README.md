# Member

用户基础服务，基于RBAC

## 运行命令

运行api

```bash
go run ./api/main.go -f ./api/etc/main.yaml
```

运行rpc

```bash
go run ./rpc/main.go -f ./rpc/etc/main.yaml
```

## 操作命令（开发时使用）

添加缺少的模块

```bash
go mod tidy
```

切换目录

```bash
cd E:\MyProjects\crucible_go\app\core\user\member
```

## 生成命令

生成Model，通过sql文件生成，数据库生成存在类型不匹配的问题

```bash
goctl model mysql ddl --src=proto/sql/schema.sql --dir=proto/model --style goZero
```

生成API

```bash
goctl api go -api ./proto/api/main.api -dir ./api --style goZero
```

生成RPC

```bash
goctl rpc protoc ./proto/rpc/main.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc -m --style goZero
```

## 配置文件

### RPC

`app/core/user/member/rpc/etc/main.yaml`

```yaml
Name: user.member.rpc
ListenOn: 127.0.0.1:8098

# MySQL
DataSource: root:123456@tcp(127.0.0.1:3306)/crucible_go?parseTime=true

# Consul:
Consul:
  Host: 127.0.0.1:8500
#  Token: ''
  Key: user.member.rpc
  Meta:
    Protocol: grpc
  Tag:
    - user.member
    - rpc
```

`app/core/user/member/rpc/internal/config/config.go`

```go
package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	// db url:
	DataSource string

	//// cache:
	//Cache cache.CacheConf

	// 服务注册/发现:
	Consul consul.Conf
}
```


`app/core/user/member/rpc/internal/svc/servicecontext.go`

```go
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
		Member: model.NewSysMemberModel(sqlx.NewMysql(c.DataSource)),
		//Member: model.NewSysMemberModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
```
Logic 查询用户详情示例

```go
// 查询用户详情
func (l *QuerySysMemberDetailLogic) QuerySysMemberDetail(in *proto.QuerySysMemberDetailReq) (*proto.QuerySysMemberDetailResp, error) {
	// todo: add your logic here and delete this line
	member := l.svcCtx.Member

	res, _ := member.FindOne(l.ctx, in.Id)

	if res == nil {
		return nil, nil
	}

	return &proto.QuerySysMemberDetailResp{
		Id:          res.Id,
		Identify:    res.Identify,
		Username:    res.Username,
		Nickname:    res.Nickname,
		Avatar:      res.Avatar,
		Quote:       res.Quote,
		Birthday:    res.Birthday.String(),
		Gender:      res.Gender,
		Email:       res.Email,
		Telephone:   res.Telephone,
		LoginAt:     res.LoginAt.String(),
		LoginIp:     res.LoginIp,
		LoginRegion: res.LoginRegion,
		LoginOs:     res.LoginOs,
		Status:      res.Status,
		Remark:      res.Remark,
		CreateAt:    res.CreateAt.String(),
		CreateBy:    res.CreateBy,
		UpdateAt:    res.UpdateAt.String(),
		UpdateBy:    res.UpdateBy,
	}, nil
}
```

`app/core/user/member/rpc/main.go`

```go
// register service to consul
_ = consul.RegisterService(c.ListenOn, c.Consul)
```

### API

`app/core/user/member/api/etc/main.yaml`

```yaml
Name: main
Host: 127.0.0.1
Port: 8888

MemberRpc:
  Target: consul://127.0.0.1:8500/user.member.rpc?wait=14s
```

`app/core/user/member/api/internal/config/config.go`

```go
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	MemberRpc zrpc.RpcClientConf
}
```

`app/core/user/member/api/internal/svc/servicecontext.go`

```go
package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"member/api/internal/config"
	"member/rpc/client/sysmemberservice"
)

type ServiceContext struct {
	Config       config.Config
	MemberClient sysmemberservice.SysMemberService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MemberClient: sysmemberservice.NewSysMemberService(zrpc.MustNewClient(c.MemberRpc)),
	}
}
```

Logic 查询用户详情示例

```go
func (l *QuerySysMemberDetailLogic) QuerySysMemberDetail(req *types.QuerySysMemberDetailReq) (resp *types.QuerySysMemberDetailResp, err error) {
	// todo: add your logic here and delete this line
	res, _ := l.svcCtx.MemberClient.QuerySysMemberDetail(l.ctx, &proto.QuerySysMemberDetailReq{
		Id: req.Id,
	})

	if res == nil {
		return nil, nil
	}

	data := types.QuerySysMemberDetailData{
		Id:          res.Id,
		Identify:    res.Identify,
		Username:    res.Username,
		Nickname:    res.Nickname,
		Avatar:      res.Avatar,
		Quote:       res.Quote,
		Birthday:    res.Birthday,
		Gender:      res.Gender,
		Email:       res.Email,
		Telephone:   res.Telephone,
		LoginAt:     res.LoginAt,
		LoginIp:     res.LoginIp,
		LoginRegion: res.LoginRegion,
		LoginOs:     res.LoginOs,
		Status:      res.Status,
		Remark:      res.Remark,
		CreateAt:    res.CreateAt,
		CreateBy:    res.CreateBy,
		UpdateAt:    res.UpdateAt,
		UpdateBy:    res.UpdateBy,
	}

	return &types.QuerySysMemberDetailResp{
		Code:    "200",
		Message: "OK",
		Data:    data,
	}, nil
}
```

`app/core/user/member/api/main.go`

```go
_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
```