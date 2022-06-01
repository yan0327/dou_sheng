# 部署指南

### Docker 
1. 构建镜像，在项目根目录执行
```bash
make image
```
2. 推送到镜像仓库（可选）
```bash
docker login <仓库>
docker tag douyin <仓库>/douyin[:标签]
docker push <仓库>/douyin[:标签]
```
3. 创建存储卷与网络
```bash
docker volume create douyin-config
docker volume create douyin-storage
docker network create douyin-net
```

4. 拷贝配置文件到 `douyin-config` volume 下，并按需修改
```bash
cp configs/* $(docker volume inspect douyin-config | grep Mountpoint | awk '{print $2}' | awk '{gsub("[,\"]", ""); print $0}')
```

5. 启动相关容器
   1. 启动 jaeger 链路追踪
    ```bash
    docker run -d --net douyin-net \
      --name douyin-jaeger \
      jaegertracing/all-in-one
    ```
   2. 启动 Minio 对象存储
    ```bash
    docker run -d --net douyin-net --name douyin-minio \
      -v douyin-storage:/data \
      minio/minio server /data
    ```
   3. 为 Minio 创建 Bucket（Bucket 名称需与配置文件一致）
   ```bash
   docker run --net douyin-net --rm -it --entrypoint=bin/sh minio/mc
   mc config host add minio http://douyin-minio:9000 minioadmin minioadmin
   mc mb minio/<Bucket名称>
   ```
   4. 启动 MySQL 数据库（这里的环境变量需与配置文件一致） 
    ```bash
    docker run -d --net douyin-net --name douyin-db \
      -e MYSQL_DATABASE=<数据库名> \
      -e MYSQL_ROOT_PASSWORD=<ROOT用户密码> \
      -e TZ=Asia/Shanghai \
      -v <项目根绝对目录>/third_party/sql:/docker-entrypoint-initdb.d \
      mysql
    ```
   5. 启动应用服务器
    ```bash
    docker run -d --net douyin-net \
      --name douyin-server \
      -p 18080:8080 \
      -v douyin-config:/app/configs \
      douyin
    ```
---

### Docker Compose

TODO

---

### Kubernetes

TODO