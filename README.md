# Crucible Go

某人半路使用 Go 重写的毕业设计💻.

基于 `go` + `go-zero` + `gRPC` + `grom`, 以`微服务`方式实现的`ACM/ICPC`在线判题平台🖥️.

## 🏠架构

🔖参考 [zeromall](https://github.com/zeromicro/zeromall) 的架构公式

📝公式 [biz](app/biz) = [core](app/core) + [unit](app/unit)

- [biz](app/biz): Business 产品
- [core](app/core): 核心的基础服务
- [unit](app/unit): 通用的业务单元

## 📂目录结构

```text
crucible_go
├── go.work         # Go 工作区文件，用于将多个模块（即微服务）连接在一起
├─app
│  ├─biz
│  ├─core
│  │  └─user
│  │      ├─auth
│  │      └─member
│  │          ├─api
│  │          ├─proto
│  │          │  ├─api
│  │          │  ├─model
│  │          │  ├─rpc
│  │          │  └─sql
│  │          └─rpc
│  └─unit
├─docs
└─script
```

## 📖参考资料

### 网站

- [oi-wiki](https://oi-wiki.org/)

### 文献

- ...