创建网络

```bash
docker network create dev-network
```

MySQL

```bash
docker run -d --name mysql-server --network dev-network --restart always -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql:latest
```

Etcd

```bash
docker run -d --name etcd-server --network dev-network --restart always -e ALLOW_NONE_AUTHENTICATION=yes -p 2379:2379 -p 2380:2380 bitnami/etcd:latest
```

Prometheus

```bash
docker run -d --name prometheus-server --network dev-network --restart always -v E:/MyProjects/crucible_go/script/prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml -p 9090:9090 bitnami/prometheus:latest
```