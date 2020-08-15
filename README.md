# HttpForward

用于http请求转发,支持http https

## 环境依赖

go v1.13.0 +

## 编译步骤

```shell
   #windows 编译
   1. set GOARCH=amd64
   2. set GOOS=linux
   3. go build .
```

## 部署步骤

```shell
   1. cd /data
   2. mkdir httpforward
   3. cd httpforward
   4. mkdir app
   5. 上传编译后的HttpForward至app目录,docker-compose.yml,Dockerfile至httpforward目录
   6. docker-compose up --build
   7. chmod -R 777 .
```
## 使用示例

http://127.0.0.1:6060?url=http://www.google.com

