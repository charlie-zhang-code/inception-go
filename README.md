## crucible_go

某人半路使用 Go 重写的毕业设计，基于 `go` + `go-zero` + `gRPC`, 以`微服务`方式实现的`ACM/ICPC`在线判题平台，企图一劳永逸.

### 架构

参考 [zeromall](https://github.com/zeromicro/zeromall) 的架构公式

公式 [biz](app/biz) = [core](app/core) + [unit](app/unit)

- [biz](app/biz):  Business 产品
- [core](app/core): 核心的基础服务
- [unit](app/unit): 通用的业务单元

工程目录

```
crucible_go
├── go.work         # Go 工作区文件，用于将多个模块（即微服务）连接在一起
├── app/
│   ├── biz
│   ├── core
│   │   └── user   # 用户相关的核心服务
│   └── unit
├── client
├── docs
├── script
└── README.md
```

之所以采用这个架构是因为想一次性写好核心服务的框架，后面新项目的时候可以直接使用核心架构.

暂时命名为 `crucible_go`，后面会改名，目前只有 `ACM/ICPC` 在线判题平台，预期自己的博客也集成在这个架构.

### 参考资料

#### 网站

- [oi-wiki](https://oi-wiki.org/)

#### 文献

- ...