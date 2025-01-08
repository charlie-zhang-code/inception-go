# Docker

## 创建网络

创建 dev-network 的 Docker 网络. 让容器连接到同一个网络上，方便地相互通信.

```bash
docker network create dev-network
```

## 构建容器

### MySQL

持久化存储

```bash
docker run -d --name mysql-server --network dev-network --restart always -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql:latest
```

### Redis

应用层的高速缓存

```bash
docker run -d --name redis-server --network dev-network --restart always -p 6379:6379 redis:latest
```

加上密码接上

```bash
... redis:latest redis-server --requirepass 123456
```

### Consul

服务自动注册和发现

```bash
docker run -d --name=consul-server --network dev-network --restart always -p 8500:8500 consul:1.15.4 agent -server -bootstrap -ui -node=default -client='0.0.0.0'
```

下面的暂时用不上

Etcd

```bash
docker run -d --name etcd-server --network dev-network --restart always -e ALLOW_NONE_AUTHENTICATION=yes -p 2379:2379 -p 2380:2380 bitnami/etcd:latest
```

Prometheus

```bash
docker run -d --name prometheus-server --network dev-network --restart always -v E:/MyProjects/crucible_go/script/prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml -p 9090:9090 bitnami/prometheus:latest
```