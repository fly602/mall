# 商城后端系统

## 下载
```
git clone git@github.com:fly602/mall.git
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
