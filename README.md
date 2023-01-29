# 商城后端系统
基于go-zero搭建的商城后端微服务。集成jwt鉴权、服务监控、链路追踪、分布式事务、日志采集等服务。

## 服务介绍

```
├── dtm                 # DTM 分布式事务管理器
├── elasticsearch       # ES 分布式全文检索引擎，用于日志采集
├── etcd                # Etcd 服务注册发现
├── filebeat            # Filebeat 用于日志采集
├── golang              # Golang 开发运行环境
├── go-stash            # Go-stash 用于日志采集
├── grafana             # Grafana 可视化数据监控
├── jaeger              # Jaeger 链路追踪
├── kafka               # Kafka 消息队列
├── kibana              # Kibana 日志采集可视化显示
├── mysql               # Mysql 服务
├── prometheus          # Prometheus 服务监控
├── redis               # Redis 服务
├── service             # Service 商城后端服务
│   ├── common          # Common 公共模块
│   │   ├── cryptx      # Cryptx 密码加密模块
│   │   └── jwt         # JWT 鉴权
│   ├── order           # Order 订单服务
│   │   ├── api         
│   │   ├── model
│   │   └── rpc
│   ├── pay             # Pay 支付订单服务
│   │   ├── api
│   │   ├── model
│   │   └── rpc
│   ├── product         # Product 产品服务
│   │   ├── api
│   │   ├── model
│   │   └── rpc
│   └── user            # User 用户服务
│       ├── api
│       ├── model
│       └── rpc
└── zookeeper
```

## 下载
```
git clone https://github.com/fly602/mall.git
```

## 编译
```sh
cd mall
## filebeat.yml需要root权限
sudo chown -R root filebeat/filebeat.yml
cd service
go mod init
cd ..
make all
```

